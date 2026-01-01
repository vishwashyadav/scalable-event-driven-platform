package processors

import (
	"api-service/internal/domain"
	"api-service/internal/service"
	"time"
)

type OrderProcessor struct {
	service *service.OrderService
}

func NewOrderProcessor(s *service.OrderService) *OrderProcessor {
	return &OrderProcessor{
		service: s,
	}
}

func (p *OrderProcessor) Process(orderId string) {
	go func() {
		p.service.UpdateOrderStatus(orderId, domain.OrderProcessing)
		time.Sleep(5 * time.Second)

		p.service.UpdateOrderStatus(orderId, domain.OrderCompleted)

	}()
}
