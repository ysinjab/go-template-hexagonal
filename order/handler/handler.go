package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ysinjab/go-template-hexagonal/order/service"
)

type OrderDetails struct {
	ProductID string `json:"product_id"`
}

type OrderHandler interface {
	OrderHandler(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	orderService service.OrderService
}

func (h *orderHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var body []*OrderDetails

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productIds := make([]string, 0)
	for _, od := range body {
		productIds = append(productIds, od.ProductID)
	}

	if len(productIds) == 0 {
		http.Error(w, "no product id provided", http.StatusBadRequest)
		return
	}

	err = h.orderService.CreateOrder(r.Context(), productIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func New(orderService service.OrderService) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}
