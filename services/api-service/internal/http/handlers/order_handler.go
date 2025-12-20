package handlers

import (
	"api-service/internal/domain"
	"api-service/internal/service"
	"api-service/internal/utility"
	"fmt"
	"net/http"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(s *service.OrderService) *OrderHandler {

	return &OrderHandler{service: s}
}

func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var eventId string
	switch r.URL.Path {
	case "/orders/CreateOrder":
		payLoad, err := utility.DecodeBody[domain.Order](w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		eventId, _ = h.service.CreateOrder(payLoad)
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// success
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fmt.Sprint(eventId)))
}
