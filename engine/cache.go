package main

import (
	"encoding/json"
	"os"

	"github.com/briandowns/stock-exchange/models"
)

const symbolDataFile = "data/nasdaq.json"

// Cache holds the implemented caching system
type Cache struct {
	Cacher
}

// Cacher represents cache behavior
type Cacher interface {
	Build() error
	Get(key []byte) (models.Company, error)
	Add(key []byte, value models.Company) error
	Entries() ([]models.Company, error)
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
