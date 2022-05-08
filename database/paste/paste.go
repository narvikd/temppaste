package paste

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"temppaste/pkg/errorskit"
	"time"
)

// Paste represents the table, it also includes validation directives.
type Paste struct {
	Id      string `json:"id"`
	Content string `json:"content" validate:"required,lte=524288"` // paste byte limit (512 * 1024 = 512kb)
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

func GetPaste(db *memdb.MemDB, id string) (*Paste, error) {
	txn := db.Txn(false) // Create a read transaction
	defer txn.Abort()    // Aborts in case of an error
	raw, err := txn.First("paste", "id", id)
	if err != nil {
		return nil, errorskit.Wrap(err, "couldn't get post posts")
	}
	if raw == nil {
		return nil, errors.New("paste not found")
	}
	return raw.(*Paste), nil
}

// NewPaste accepts a Paste pointer to create a new paste in the DB.
func NewPaste(db *memdb.MemDB, paste *Paste) (string, error) {
	const pasteDelTime = 5 * time.Minute
	id := uuid.New().String()
	paste.Id = id // Sets the ID

	txn := db.Txn(true) // Create a read transaction
	err := txn.Insert("paste", paste)
	if err != nil {
		return id, errorskit.Wrap(err, "couldn't insert paste")
	}
	txn.Commit()

	deletePasteAfterTime(db, paste, pasteDelTime)
	return id, nil
}

// deletePasteAfterTime wraps DeletePaste in a go func to delete the just created paste after X time.
func deletePasteAfterTime(db *memdb.MemDB, paste *Paste, duration time.Duration) {
	go func() {
		time.Sleep(duration)
		errDelete := deletePaste(db, paste)
		if errDelete != nil {
			errorskit.LogWrap(errDelete, "couldn't delete paste")
		}
	}()
}

// deletePaste deletes the Paste passed as pointer.
func deletePaste(db *memdb.MemDB, paste *Paste) error {
	txn := db.Txn(true) // Create a read transaction
	err := txn.Delete("paste", paste)
	if err != nil {
		return errorskit.Wrap(err, "couldn't delete paste")
	}
	txn.Commit()
	return nil
}
