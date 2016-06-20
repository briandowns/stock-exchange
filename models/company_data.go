package models

// Company represents a company
type Company struct {
	Symbol       string  `json:"Symbol"`
	Name         string  `json:"Name"`
	LastSale     float64 `json:"LastSale"`
	MarketCap    float64 `json:"MarketCap"`
	IPOYear      int     `json:"IPOyear"`
	Sector       string  `json:"Sector"`
	Industry     string  `json:"industry"`
	SummaryQuote string  `json:"SummaryQuote"`
}
