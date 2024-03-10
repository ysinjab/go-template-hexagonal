package domain

import (
	crand "crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

type Payment struct {
	ID         string
	CustomerID string
	Status     string
	OrderID    string
	TotalPrice float64
	CreatedAt  time.Time
}

func NewPayment(customerID, orderID string, totalPrice float64) *Payment {
	return &Payment{
		ID:         ulid.MustNew(ulid.Now(), crand.Reader).String(),
		CustomerID: customerID,
		Status:     "PENDING",
		OrderID:    orderID,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
	}
}
