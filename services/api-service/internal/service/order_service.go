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

func (s *OrderService) GetOrderById(orderId string) (*domain.Order, error) {
	return s.repository.FindById(orderId)
}

func (s *OrderService) GetAll() ([]*domain.Order, error) {
	return s.repository.GetAll()
}

func (s *OrderService) UpdateOrderStatus(orderId string, status domain.OrderStatus) error {
	order, err := s.repository.FindById(orderId)
	if err != nil {
		return err
	}
	if !s.canTransition(order.Status, status) {
		return errors.New("invalid status transition")
	}
	return s.repository.UpdateStatus(orderId, status)
}

func (s *OrderService) canTransition(current, new domain.OrderStatus) bool {
	switch current {
	case domain.OrderCreated:
		return new == domain.OrderProcessing
	case domain.OrderProcessing:
		return new == domain.OrderCompleted
	case domain.OrderCompleted:
		return false
	}
	return false
}
