package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	authv1 "github.com/petrrrnikitin/flagctl/services/auth/gen"
	"github.com/petrrrnikitin/flagctl/services/auth/internal/handler"
)

func main() {
	addr := os.Getenv("AUTH_GRPC_ADDR")
	if addr == "" {
		addr = ":50051"
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	authv1.RegisterAuthServiceServer(srv, &handler.AuthHandler{})

	log.Printf("auth service listening on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}