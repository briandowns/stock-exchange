package main

import (
	"encoding/json"
	"os"

	"github.com/briandowns/stock-exchange/models"
)

const companyDataFile = "data/companylist.csv"

// generateCompanyData
func generateCompanyData() (models.CompanyData, error) {
	f, err := os.Open(companyDataFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cd models.CompanyData
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&cd); err != nil {
		return nil, err
	}
	return cd, nil
}
