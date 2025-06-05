package handler

import (
	"net/http"

	"github.com/javy99/kitchen/services/common/genproto/orders"
	"github.com/javy99/kitchen/services/common/util"
	"github.com/javy99/kitchen/services/orders/types"
)

type OrderHttpHandler struct {
	orderService types.OrderService
}

func NewHttpOrdersServer(ordersService types.OrderService) *OrderHttpHandler {
	handler := &OrderHttpHandler{
		orderService: ordersService,
	}

	return handler
}

func (h *OrderHttpHandler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrderHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
