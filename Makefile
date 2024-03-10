# Go parameters
GOCMD=go
BINARY_UNIX=$(BINARY_NAME)_unix

# Protobuf parameters
PROTOC=protoc
PROTO_DIR=./proto
PROTO_GEN_DIR=./genproto


# Generate gRPC code from .proto files
protogen:
	mkdir -p $(PROTO_GEN_DIR)
	$(PROTOC) -I$(PROTO_DIR) --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative \
	          --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative \
	          $(PROTO_DIR)/payment/v1/payment.proto



