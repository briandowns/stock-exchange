package database

import (
	"github.com/boltdb/bolt"
)

// NewDB
func NewDB() (*bolt.DB, error) {
	return bolt.Open("blog.db", 0600, nil)
}
