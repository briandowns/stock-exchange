package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/briandowns/stock-exchange/models"
)

var errUnknownCache = errors.New("unknown cache type")

// Cache holds the implemented caching system
type Cache struct {
	Cacher
}

// Cacher represents cache behavior
type Cacher interface {
	Build() error
	Get(key []byte) (models.Company, error)
	Add(key []byte, value models.Company) error
	//Entries() ([]models.Company, error)
	Entries() ([]models.Company, error)
}

// generateSymbolData
func generateSymbolData(symbolFile string) ([]models.Company, error) {
	f, err := os.Open(symbolFile)
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
