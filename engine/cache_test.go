package main

import "testing"

var configFile = "../config.json"
var exchanges = []string{"nasdaq", "nyse", "amex"}

// generateSymbolData
func TestGenerateSymbolData(t *testing.T) {
	for _, exchange := range exchanges {
		_, err := generateSymbolData("data/" + exchange + "_symbols.json")
		if err != nil {
			t.Error(err)
		}
	}
}
