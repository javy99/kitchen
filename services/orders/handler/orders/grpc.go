package handler

import (
	"context"

	"github.com/javy99/kitchen/services/common/genproto/orders"
	"github.com/javy99/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	// service injection
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGRPCOrdersServer(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrderGrpcHandler{
		orderService: ordersService,
	}

	// register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}
	return res, nil
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
