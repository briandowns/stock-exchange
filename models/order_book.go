package models

import (
	"sort"
	"sync"
	"time"

	"gopkg.in/validator.v2"
)

// Order represents an order
type Order struct {
	ID       string    `validate:len=32`
	Symbol   string    `validate:"nonzero"`
	Time     time.Time `validate:"nonzero"`
	Quantity int       `validate:"nonzero"`
	Bid      float64   `validate:"nonzero"`
}

// Validate will make sure that all fields are filled in
func (o Order) Validate() bool {
	if errs := validator.Validate(o); errs != nil {
		return false
	}
	return true
}

// OrderBook represents the current state of the market
type OrderBook struct {
	Lock   sync.Locker `json:"-"`
	Orders []Order     `json:"orders"`
}

// Len gets the length of the
func (o *OrderBook) Len() int {
	return len(o.Orders)
}

// Less determins if the first value is before the second
func (o *OrderBook) Less(i, j int) bool {
	return o.Orders[i].Time.Before(o.Orders[j].Time)
}

// Swap will swap the placement of the elements by their indexes
func (o *OrderBook) Swap(i, j int) {
	o.Orders[i], o.Orders[j] = o.Orders[j], o.Orders[i]
}

// NewOrderBook creates a new value of type OrderBook pointer
func NewOrderBook() *OrderBook {
	return &OrderBook{
		Lock:   &sync.Mutex{},
		Orders: make([]Order, 0),
	}
}

// Add adds an order to the book
func (o *OrderBook) Add(order Order) error {
	o.Lock.Lock()
	defer o.Lock.Unlock()
	o.Orders = append(o.Orders, order)
	sort.Sort(o)
	return nil
}

// Cancel removes an order from the book
func (o *OrderBook) Cancel(order Order) error {
	o.Lock.Lock()
	defer o.Lock.Unlock()
	for idx, i := range o.Orders {
		if i.ID == order.ID {
			o.Orders = append(o.Orders[:1], o.Orders[idx+1:]...)
			return nil
		}
	}
	return nil
}

// Execute
func (o *OrderBook) Execute() {}
