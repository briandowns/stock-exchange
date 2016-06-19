package models

import "fmt"

// Company holds the csv unencoded data from the companies file that
// needs to be parsed. Ex. below
//"Symbol","Name","LastSale","MarketCap","IPOyear","Sector","industry","Summary Quote",
type Company struct {
	Symbol       string
	Name         string
	LastSale     string
	MarketCap    string
	IPYear       string
	Sector       string
	Industry     string
	SummaryQuote string
}

// NewCompany
func NewCompany(data []string) Company {
	var c Company
	c.Symbol = data[0]
	c.Name = data[1]
	c.LastSale = data[2]
	c.MarketCap = data[3]
	c.IPYear = data[4]
	c.Sector = data[5]
	c.Industry = data[6]
	c.SummaryQuote = data[7]
	fmt.Println(c)
	return c
}

// CompanyData contains all company data
type CompanyData []Company

func NewCompanyData() CompanyData {
	cd := make([]Company, 0)
	return cd
}
