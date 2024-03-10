package domain

import (
	crand "crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

type Order struct {
	ID           string
	CustomerID   string
	OrderedAt    time.Time
	OrderDetails []*OrderDetails
}

func (o *Order) CalculateTotalPrice() float64 {
	var totalPrice float64
	for _, od := range o.OrderDetails {
		totalPrice += od.UnitPrice
	}
	return totalPrice
}

func NewOrder(customerID string) *Order {
	return &Order{
		ID:         ulid.MustNew(ulid.Now(), crand.Reader).String(),
		CustomerID: customerID,
		OrderedAt:  time.Now(),
	}
}
