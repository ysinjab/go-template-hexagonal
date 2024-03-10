package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	paymentpbv1 "github.com/ysinjab/go-template-hexagonal/genproto/payment/v1"
	"github.com/ysinjab/go-template-hexagonal/payment/repository"
	"github.com/ysinjab/go-template-hexagonal/payment/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	pool, err := pgxpool.New(context.Background(), "postgres://localhost:5432/payments")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	paymentpbv1.RegisterPaymentServiceServer(s, service.NewServer(repository.New(pool)))

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
