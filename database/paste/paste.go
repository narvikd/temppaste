package paste

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"log"
	"temppaste/pkg/errorskit"
	"time"
)

// Paste represents the table, it also includes validation directives.
type Paste struct {
	Id      string
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func NewSchema() *memdb.DBSchema {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"paste": {
				Name: "paste",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
				},
			},
		},
	}
	return schema
}

// GetAllPastes returns a slice with all the DB's pastes.
func GetAllPastes(db *memdb.MemDB) ([]Paste, error) {
	var pastes []Paste

	txn := db.Txn(false) // Create a read transaction
	defer txn.Abort()

	resultIterator, err := txn.Get("paste", "id")
	if err != nil {
		return nil, errorskit.Wrap(err, "couldn't get all posts")
	}

	for obj := resultIterator.Next(); obj != nil; obj = resultIterator.Next() {
		o := obj.(*Paste)
		paste := Paste{
			Id:      o.Id,
			Name:    o.Name,
			Content: o.Content,
		}
		pastes = append(pastes, paste)
	}

	return pastes, nil
}

// NewPaste accepts a Paste pointer to create a new paste in the DB.
func NewPaste(db *memdb.MemDB, paste *Paste) error {
	const pasteDelTime = 5 * time.Minute
	paste.Id = uuid.New().String() // Sets the ID

	txn := db.Txn(true) // Create a read transaction
	err := txn.Insert("paste", paste)
	if err != nil {
		return errorskit.Wrap(err, "couldn't insert paste")
	}
	txn.Commit()

	deletePasteAfterTime(db, paste, pasteDelTime)
	return nil
}

// deletePasteAfterTime wraps DeletePaste in a go func to delete the just created paste after X time.
func deletePasteAfterTime(db *memdb.MemDB, paste *Paste, duration time.Duration) {
	go func() {
		time.Sleep(duration)
		errDelete := DeletePaste(db, paste)
		if errDelete != nil {
			log.Println(errorskit.Wrap(errDelete, "couldn't delete paste"))
		}
	}()
}

// DeletePaste deletes a Paste passed as pointer.
func DeletePaste(db *memdb.MemDB, paste *Paste) error {
	txn := db.Txn(true) // Create a read transaction
	err := txn.Delete("paste", paste)
	if err != nil {
		return errorskit.Wrap(err, "couldn't delete paste")
	}
	txn.Commit()
	return nil
}
