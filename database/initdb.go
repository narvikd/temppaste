package database

import (
	"github.com/hashicorp/go-memdb"
	"temppaste/pkg/errorskit"
)

// NewDB returns a new "go-memdb" pointer using the "Paste" schema.
func NewDB() (*memdb.MemDB, error) {
	db, err := memdb.NewMemDB(newSchema())
	if err != nil {
		return nil, errorskit.Wrap(err, "couldn't create new DB")
	}
	return db, nil
}

func newSchema() *memdb.DBSchema {
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
