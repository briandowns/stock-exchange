package database

import (
	"github.com/boltdb/bolt"
)

// NewDB is called when a new database is needed
func NewDB(dbName string) (*bolt.DB, error) {
	return bolt.Open(dbName, 0600, nil)
}
