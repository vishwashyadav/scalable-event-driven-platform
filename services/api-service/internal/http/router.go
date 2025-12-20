package router

import (
	"api-service/internal/http/handlers"
	"api-service/internal/repository"
	"api-service/internal/service"
	"net/http"
)

type Router struct{}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	orderRep := repository.OrderRepository{}
	orderService := service.NewOrderService(&orderRep)
	handler := handlers.NewOrderHandler(orderService)
	// Register both the exact path and the prefix with a trailing slash.
	// The ServeMux treats patterns not ending with "/" as exact matches,
	// so to match subpaths like "/orders/CreateOrder" we must register
	// the prefix with a trailing slash.
	mux.Handle("/orders/", handler)
	return mux
}
