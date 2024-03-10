package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ysinjab/go-template-hexagonal/order/domain"
)

type Repository interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	CreateOrderDetails(ctx context.Context, orderDetails *domain.OrderDetails) error
}

type repository struct {
	db *pgxpool.Pool
}

func (r *repository) CreateOrder(ctx context.Context, o *domain.Order) error {
	sql := `INSERT INTO orders (id, customer_id, ordered_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, sql, o.ID, o.CustomerID, o.OrderedAt)
	if err != nil {
		return fmt.Errorf("failed to insert order: %v", err)
	}

	return nil
}

func (r *repository) CreateOrderDetails(ctx context.Context, od *domain.OrderDetails) error {
	sql := `INSERT INTO order_details (id, order_id, product_id, unit_price) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, sql, od.ID, od.OrderID, od.ProductID, od.UnitPrice)
	if err != nil {
		return fmt.Errorf("failed to insert order details: %v", err)
	}

	return nil
}

func New(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}
