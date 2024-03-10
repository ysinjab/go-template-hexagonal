package service

import (
	"context"
	"log"

	paymentpbv1 "github.com/ysinjab/go-template-hexagonal/genproto/payment/v1"
	"github.com/ysinjab/go-template-hexagonal/order/domain"
	orderrepository "github.com/ysinjab/go-template-hexagonal/order/repository"
)

type OrderService interface {
	CreateOrder(context.Context, []string) error
}

type service struct {
	repository orderrepository.Repository
	payment    paymentpbv1.PaymentServiceClient
}

func (s *service) CreateOrder(ctx context.Context, productIds []string) error {
	// TODO: get customer id from context
	o := domain.NewOrder("X")
	o.OrderDetails = make([]*domain.OrderDetails, 0)
	for _, pid := range productIds {
		o.OrderDetails = append(o.OrderDetails, domain.NewOrderDetails(o.ID, pid, 1.0))
	}

	// TODO: those should be in a transaction
	if err := s.repository.CreateOrder(ctx, o); err != nil {
		return err
	}

	for _, od := range o.OrderDetails {
		if err := s.repository.CreateOrderDetails(ctx, od); err != nil {
			return err
		}
	}

	_, err := s.payment.CreatePayment(ctx, &paymentpbv1.CreatePaymentRequest{
		OrderId:    o.ID,
		CustomerId: o.CustomerID,
		TotalPrice: o.CalculateTotalPrice(),
	})
	if err != nil {
		log.Printf("failed to create payment: %v", err)
		return err
	}

	return nil
}

func New(repository orderrepository.Repository, payment paymentpbv1.PaymentServiceClient) OrderService {
	return &service{
		repository: repository,
		payment:    payment,
	}
}
