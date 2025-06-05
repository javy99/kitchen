package main

import (
	"log"
	"net/http"

	handler "github.com/javy99/kitchen/services/orders/handler/orders"
	"github.com/javy99/kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersServer(orderService)
	orderHandler.RegisterRoute(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
