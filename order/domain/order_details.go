package domain

import (
	crand "crypto/rand"

	"github.com/oklog/ulid"
)

type OrderDetails struct {
	ID        string
	OrderID   string
	ProductID string
	UnitPrice float64
}

func NewOrderDetails(orderID, productID string, unitPrice float64) *OrderDetails {
	return &OrderDetails{
		ID:        ulid.MustNew(ulid.Now(), crand.Reader).String(),
		ProductID: productID,
		OrderID:   orderID,
		UnitPrice: unitPrice,
	}
}
