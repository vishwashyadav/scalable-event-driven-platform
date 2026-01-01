package handlers

import (
	"api-service/internal/domain"
	"api-service/internal/processors"
	"api-service/internal/service"
	"api-service/internal/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

type OrderHandler struct {
	service   *service.OrderService
	processor *processors.OrderProcessor
}

func NewOrderHandler(s *service.OrderService, p *processors.OrderProcessor) *OrderHandler {

	return &OrderHandler{service: s, processor: p}
}

func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := strings.TrimSpace(path.Clean(r.URL.Path))
	fmt.Println("Handling path:", p)
	var eventId string
	switch p {
	case "/orders/CreateOrder":
		payLoad, _ := utility.DecodeBody[domain.Order](w, r)
		eventId, _ = h.service.CreateOrder(payLoad)
		h.processor.Process(eventId)
		// success
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(fmt.Sprint(eventId)))
	case "/orders/GetById":
		// validate id parameter
		id := r.URL.Query().Get("id")
		if strings.TrimSpace(id) == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}
		order, _ := h.service.GetOrderById(id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
		return
	case "/orders/GetAll":
		orders, _ := h.service.GetAll()
		w.WriteHeader((http.StatusOK))
		json.NewEncoder(w).Encode(orders)
		return
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

}
