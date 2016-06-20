package database

import (
	"github.com/boltdb/bolt"
)

// NewDB is called when a new database is needed
func NewDB() (*bolt.DB, error) {
	return bolt.Open("blog.db", 0600, nil)
}
