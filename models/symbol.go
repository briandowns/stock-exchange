package models

import (
	"regexp"
)

// CUSIP represents a stock CUSIP
type CUSIP string

// Validate will validate a CUSIP
func (c CUSIP) Validate() bool {
	r := regexp.MustCompile(`[0-9]{3}[a-zA-Z0-9]{6}`)
	if r.MatchString(string(c)) {
		return false
	}
	return true
}

// Symbol represents a stock
type Symbol struct {
	CUSIP
	Name string
}
