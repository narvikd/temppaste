package database

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"log"
	"temppaste/pkg/errorskit"
	"time"
)

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

// NewPaste accepts a Paste pointer to create a new paste.
func NewPaste(db *memdb.MemDB, paste *Paste) error {
	paste.Id = uuid.New().String() // Sets the ID

	txn := db.Txn(true) // Create a read transaction
	err := txn.Insert("paste", paste)
	if err != nil {
		return errorskit.Wrap(err, "couldn't insert paste")
	}
	txn.Commit()

	deletePasteAfterTime(db, paste, 5)
	return nil
}

// deletePasteAfterTime wraps DeletePaste in a go func to delete the paste after X minutes.
func deletePasteAfterTime(db *memdb.MemDB, paste *Paste, minutes time.Duration) {
	go func() {
		time.Sleep(minutes * time.Minute)
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
