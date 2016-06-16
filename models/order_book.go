package models

import (
	"time"
)

type Order struct {
	Time     time.Time
	Quantity int
	Bid      float64
}

// OrderBook represents the current state of the market
type OrderBook struct {
	Orders []Order
}
