package service

import (
	"context"

	paymentpbv1 "github.com/ysinjab/go-template-hexagonal/genproto/payment/v1"
	"github.com/ysinjab/go-template-hexagonal/payment/domain"
	paymentrepo "github.com/ysinjab/go-template-hexagonal/payment/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	repository paymentrepo.Repository
	paymentpbv1.UnimplementedPaymentServiceServer
}

func (s *server) CreatePayment(ctx context.Context, req *paymentpbv1.CreatePaymentRequest) (*paymentpbv1.CreatePaymentResponse, error) {
	domainPayment := domain.NewPayment(req.CustomerId, req.OrderId, req.TotalPrice)
	if err := s.repository.Create(ctx, domainPayment); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentpbv1.CreatePaymentResponse{}, nil
}

func NewServer(repository paymentrepo.Repository) paymentpbv1.PaymentServiceServer {
	return &server{repository: repository}
}
