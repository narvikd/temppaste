package database

import (
	"github.com/hashicorp/go-memdb"
	"github.com/narvikd/errorskit"
)

// NewDB accepts a DB schema and returns a new DB object.
func NewDB(schema *memdb.DBSchema) (*memdb.MemDB, error) {
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, errorskit.Wrap(err, "couldn't create new DB")
	}
	return db, nil
}
