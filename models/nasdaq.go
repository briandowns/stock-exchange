package models

import (
	"sort"
	"sync"
)

// NasdaqOrderBook represents the current state of the market
type NasdaqOrderBook struct {
	Lock   sync.Locker `json:"-"`
	Orders []Order     `json:"orders"`
}

// Len gets the length of the
func (n *NasdaqOrderBook) Len() int {
	return len(n.Orders)
}

// Less determins if the first value is before the second
func (n *NasdaqOrderBook) Less(i, j int) bool {
	return n.Orders[i].Time.Before(n.Orders[j].Time)
}

// Swap will swap the placement of the elements by their indexes
func (n *NasdaqOrderBook) Swap(i, j int) {
	n.Orders[i], n.Orders[j] = n.Orders[j], n.Orders[i]
}

// NewNasdaqOrderBook creates a new value of type OrderBook pointer
func NewNasdaqOrderBook() *NasdaqOrderBook {
	return &NasdaqOrderBook{
		Lock:   &sync.Mutex{},
		Orders: make([]Order, 0),
	}
}

// Add adds an order to the book and sorts the book
// based on time of entry
func (n *NasdaqOrderBook) Add(order Order) error {
	n.Lock.Lock()
	defer n.Lock.Unlock()
	n.Orders = append(n.Orders, order)
	sort.Sort(n)
	return nil
}

// Cancel removes an order from the book
func (n *NasdaqOrderBook) Cancel(orderID string) error {
	n.Lock.Lock()
	defer n.Lock.Unlock()
	for idx, i := range n.Orders {
		if i.ID == orderID {
			n.Orders = append(n.Orders[:idx], n.Orders[idx+1:]...)
			return nil
		}
	}
	return nil
}

// Execute
func (n *NasdaqOrderBook) Execute() error {
	return nil
}
