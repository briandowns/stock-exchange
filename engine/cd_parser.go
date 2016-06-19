package main

import (
	"encoding/csv"
	"fmt"
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
	c := csv.NewReader(f)
	lines, err := c.ReadAll()
	if err != nil {
		return nil, err
	}
	// setting the lenfth of the slice should be more scientific
	// probably better to create a function to get the total number
	// of lines in the file minus the header and set to that.
	//	cd := make(models.CompanyData, 1000)
	var cd models.CompanyData
	for _, line := range lines {
		company := models.NewCompany(line)
		cd = append(cd, company)
	}
	fmt.Println(cd)
	return cd, nil
}
