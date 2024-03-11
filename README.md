# Hexagonal Architecture Go Project

This project follows the principles of Hexagonal Architecture, also known as Ports and Adapters, to build a modular and flexible application.

## Overview

The Hexagonal Architecture promotes a clear separation of concerns by dividing the application into three main layers:

1. **Domain Layer**: Contains the core business logic and entities of the application. It is independent of any external dependencies and frameworks.

2. **Application Layer**: Implements the use cases of the application by orchestrating the interactions between the Domain Layer and the Infrastructure Layer.


## Project Idea

The project simulates a simple e-commerce application that allows users to create and manage their orders. The application simulates payment as well. It has the following two services: orders and payment. The orders service is responsible for managing orders, and the payment service is responsible for processing payments. The two services communicate with each other through a simple gRPC call.

### Order
This is a simple RESTful http service that allows users to create and order. It has the following endpoint:
```
curl -X POST -d '[{"product_id":"1"}, {"product_id": "2"}]' http://localhost:8081/v1/orders
```

### Payment
This is a simple gRPC service that allows users to process payments. It has the following RPC method: `CreatePayment`. The payment service is called by the order service to process payments.


## Project Structure & Hexagonal Architecture
### Order
For the order service the driving port is `OrderService` under `order/serivice/service.go`. The adpter is an impleemntation of this interface. The application that use this adapter is the `OrderHandler` under `order/handler/handler.go`. The benefit here is that the `OrderHandler` is not aware of the implementation of the `OrderService`. So we can easily create a new application such as a gRPC server that uses the same `OrderService`.

The driven port for accessing the database is `Repository` under `order/repostiory/repository.go`. The adapter is an implementation of this interface that use `pgx` to store data inside PosrgreSQL. 

We have another driven port for payment service which is `PaymentServiceClient` under `genproto/payment/v1/payment_grpc.pb.go`. The adapter which is the implementation of the `PaymentServiceClient` interface lives in the same file. The adapter allows for the order service to communicate with the payment service.

### Payment
The payment service is a simple gRPC service. The driving port is the `PaymentServiceServer` interface under `genproto/payment/v1/payment_grpc.pb.go`. The adapter implementation lives under `payment/service/service.go`. 

The driven port for accessing the database is `Repository` under `payment/repostiory/repository.go`. The adapter is an implementation of this interface that use `pgx` to store data inside PosrgreSQL. 

## Running the Application
you just need to create two databases locally by run the following commands:
```
createdb payments
export POSTGRESQL_URL='postgres://localhost:5432/payments?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path payment/db/migrations up
```
```
createdb orders
export POSTGRESQL_URL='postgres://localhost:5432/orders?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path order/db/migrations up
```

then run these two commands in separate processes:
```
go run cmd/payment/main.go
```
```
go run cmd/order/main.go
```