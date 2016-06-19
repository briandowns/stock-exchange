package models

import (
	"fmt"
	"strconv"
)

// Company holds the csv unencoded data from the companies file that
// needs to be parsed. Ex. below
//"Symbol","Name","LastSale","MarketCap","IPOyear","Sector","industry","Summary Quote",
type Company struct {
	Symbol       string  `json:"Symbol"`
	Name         string  `json:"Name"`
	LastSale     float64 `json:"LastSale"`
	MarketCap    string  `json:"MarketCap"`
	IPOYear      int     `json:"IPOyear"`
	Sector       string  `json:"Sector"`
	Industry     string  `json:"industry"`
	SummaryQuote string  `json:"SummaryQuote"`
}

// NewCompany
func NewCompany(data []string) (Company, error) {
	var c Company
	var err error
	c.Symbol = data[0]
	c.Name = data[1]
	c.LastSale, err = strconv.ParseFloat(data[2], 64)
	if err != nil {
		return Company{}, nil
	}
	c.MarketCap = data[3]
	c.IPOYear, err = strconv.Atoi(data[4])
	if err != nil {
		return Company{}, nil
	}
	c.Sector = data[5]
	c.Industry = data[6]
	c.SummaryQuote = data[7]
	fmt.Println(c)
	return c, nil
}

// CompanyData contains all company data
type CompanyData []Company

func NewCompanyData() CompanyData {
	cd := make([]Company, 0)
	return cd
}
