package models

import (
	"time"

	"gopkg.in/validator.v2"
)

// OrderBooker defines the behavior of the order book
type OrderBooker interface {
	Add(order Order) error
	Cancel(orderID string) error
	Execute() error
}

// OrderBook
type OrderBook struct {
	OrderBooker
}

// Order represents an order
type Order struct {
	ID       string    `json:"id",validate:len=32`
	Symbol   string    `json:"symbol",validate:"nonzero"`
	Time     time.Time `json:"time",validate:"nonzero"`
	Quantity int       `json:"quantity",validate:"nonzero"`
	Bid      float64   `json:"bid",validate:"nonzero"`
}

// Validate will make sure that all fields are filled in
func (o Order) Validate() bool {
	if errs := validator.Validate(o); errs != nil {
		return false
	}
	return true
}
