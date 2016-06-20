package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/briandowns/stock-exchange/models"
)

const symbolCacheDB = "symbol_cache.db"
const symbolDataFile = "data/nasdaq.json"

var bucketName = []byte("symbol_cache")

// Cacher represents cache behavior
type Cacher interface {
	Build() error
	Get(key []byte) (models.Company, error)
	Put(key []byte, value models.Company) error
	All() ([]models.Company, error)
}

// SymbolCache holds the db conn
type SymbolCache struct {
	Lock sync.Locker // synchronize access to this data
	DB   *bolt.DB
}

// NewSymbolCache creates a new symbol cache
func NewSymbolCache() (*SymbolCache, error) {
	var err error
	s := SymbolCache{
		Lock: &sync.Mutex{},
	}
	s.DB, err = bolt.Open(symbolCacheDB, 0644, nil)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// generateSymbolData
func generateSymbolData() ([]models.Company, error) {
	f, err := os.Open(symbolDataFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cd []models.Company
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&cd); err != nil {
		return nil, err
	}
	return cd, nil
}

// Initialize will build the symbol cache
func (s *SymbolCache) Build() error {
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
			bucket, err := tx.CreateBucketIfNotExists(bucketName)
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
func (s *SymbolCache) Get(key []byte) (models.Company, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	var company models.Company
	err := s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
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

// Put adds a value to the cache
func (s *SymbolCache) Put(key []byte, value models.Company) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	return nil
}

// All gets all keys
func (s *SymbolCache) All() ([]models.Company, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	var cd []models.Company
	err := s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
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
