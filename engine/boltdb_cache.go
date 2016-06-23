package main

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"sync"

	"github.com/briandowns/stock-exchange/config"
	"github.com/briandowns/stock-exchange/models"

	"github.com/boltdb/bolt"
)

// BoltCache holds the db conn
type BoltCache struct {
	Lock   sync.Locker // synchronize access to this data
	DB     *bolt.DB
	config *config.Config
}

// NewBoltCache creates a new symbol cache
func NewBoltCache(db *bolt.DB, config *config.Config) *BoltCache {
	return &BoltCache{
		Lock:   &sync.Mutex{},
		DB:     db,
		config: config,
	}
}

// Build will build the symbol cache
func (s *BoltCache) Build() error {
	log.Print("Building symbol cache...")
	s.Lock.Lock()
	defer s.Lock.Unlock()

	cache, err := generateSymbolData()
	if err != nil {
		return err
	}
	for _, symbol := range cache {
		b, err := json.Marshal(symbol)
		if err != nil {
			return err
		}
		err = s.DB.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte(s.config.Cache.BoltDB.Bucket))
			if err != nil {
				return err
			}
			err = bucket.Put([]byte(symbol.Symbol), b)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	log.Println("Building symbol cache complete!")
	return nil
}

// Get retrieves a value from the cache
func (s *BoltCache) Get(key []byte) (models.Company, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	var company models.Company
	err := s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(s.config.Cache.BoltDB.Bucket))
		if bucket == nil {
			return errors.New("bucket not found")
		}
		val := bucket.Get(key)
		decoder := json.NewDecoder(strings.NewReader(string(val)))
		if err := decoder.Decode(&company); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return models.Company{}, err
	}
	return company, nil
}

// Add adds a value to the cache
func (s *BoltCache) Add(key []byte, value models.Company) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	return nil
}

// Entries gets all entries in the cache
func (s *BoltCache) Entries() ([]models.Company, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	var cd []models.Company
	err := s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(s.config.Cache.BoltDB.Bucket))
		var company models.Company
		if err := bucket.ForEach(func(k, v []byte) error {
			decoder := json.NewDecoder(strings.NewReader(string(v)))
			if err := decoder.Decode(&company); err != nil {
				return err
			}
			cd = append(cd, company)
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cd, nil
}
