package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ysinjab/go-template-hexagonal/payment/domain"
)

// there could be a domain model for payment here domain payment will be used instead

type Repository interface {
	Create(context.Context, *domain.Payment) error
}

type repository struct {
	db *pgxpool.Pool
}

func (r *repository) Create(ctx context.Context, p *domain.Payment) error {
	sql := `INSERT INTO payment (id, customer_id, status, order_id, total_price, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, sql, p.ID, p.CustomerID, p.Status, p.OrderID, p.TotalPrice, p.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert payment: %v", err)
	}

	return nil
}

func New(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}
