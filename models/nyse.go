package models

import (
	"sort"
	"sync"
)

// NYSEOrderBook represents the current state of the market
type NYSEOrderBook struct {
	Lock   sync.Locker `json:"-"`
	Orders []Order     `json:"orders"`
}

// Len gets the length of the
func (n *NYSEOrderBook) Len() int {
	return len(n.Orders)
}

// Less determins if the first value is before the second
func (n *NYSEOrderBook) Less(i, j int) bool {
	return n.Orders[i].Time.Before(n.Orders[j].Time)
}

// Swap will swap the placement of the elements by their indexes
func (n *NYSEOrderBook) Swap(i, j int) {
	n.Orders[i], n.Orders[j] = n.Orders[j], n.Orders[i]
}

// NewNYSEOrderBook creates a new value of type OrderBook pointer
func NewNYSEOrderBook() *NYSEOrderBook {
	return &NYSEOrderBook{
		Lock:   &sync.Mutex{},
		Orders: make([]Order, 0),
	}
}

// Add adds an order to the book and sorts the book
// based on time of entry
func (n *NYSEOrderBook) Add(order Order) error {
	n.Lock.Lock()
	defer n.Lock.Unlock()
	n.Orders = append(n.Orders, order)
	sort.Sort(n)
	return nil
}

// Cancel removes an order from the book
func (n *NYSEOrderBook) Cancel(orderID string) error {
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
func (n *NYSEOrderBook) Execute() error {
	return nil
}
