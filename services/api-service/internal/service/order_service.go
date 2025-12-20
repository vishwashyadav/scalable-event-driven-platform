package service

import (
	"api-service/internal/domain"
	"api-service/internal/repository"
	"errors"

	"github.com/google/uuid"
)

type OrderService struct {
	repository *repository.OrderRepository
}

func NewOrderService(r *repository.OrderRepository) *OrderService {
	return &OrderService{
		repository: r,
	}
}

func (s *OrderService) CreateOrder(order *domain.Order) (string, error) {
	if order == nil {
		return "", errors.New("amount must be greater than zero")
	}
	order.OrderId = uuid.NewString()
	_ = s.repository.Save(order)
	return order.OrderId, nil
}
