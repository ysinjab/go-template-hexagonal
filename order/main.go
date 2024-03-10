package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	paymentpbv1 "github.com/ysinjab/go-template-hexagonal/genproto/payment/v1"
	"github.com/ysinjab/go-template-hexagonal/order/handler"
	"github.com/ysinjab/go-template-hexagonal/order/repository"
	"github.com/ysinjab/go-template-hexagonal/order/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	pool, err := pgxpool.New(context.Background(), "postgres://localhost:5432/orders")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	repo := repository.New(pool)

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	payment := paymentpbv1.NewPaymentServiceClient(conn)

	r := chi.NewRouter()

	orderhandler := handler.New(service.New(repo, payment))

	r.Post("/v1/orders", orderhandler.HandleCreateOrder)

	// Start the server
	log.Printf("server listening at %v", "localhost:8081")
	http.ListenAndServe(":8081", r)
}

func adminRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return r
}
