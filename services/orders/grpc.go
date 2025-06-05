package main

import (
	"log"
	"net"

	handler "github.com/javy99/kitchen/services/orders/handler/orders"
	"github.com/javy99/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gRPCServer := grpc.NewServer()

	// register our grpc services
	orderService := service.NewOrderService()
	handler.NewGRPCOrdersServer(gRPCServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return gRPCServer.Serve(lis)
}
