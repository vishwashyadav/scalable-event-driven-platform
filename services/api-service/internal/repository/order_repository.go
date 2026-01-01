package repository

import (
	"api-service/internal/domain"
	"fmt"
	"sync"
	"time"
)

type OrderRepository struct {
	mu     sync.RWMutex
	orders map[string]*domain.Order
}

func (p *OrderRepository) Save(order *domain.Order) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.orders == nil {
		p.orders = make(map[string]*domain.Order)
	}

	if v, ok := p.orders[order.OrderId]; ok {
		return fmt.Errorf("order with id %s already exists", v.OrderId)
	}
	order.CreateAt = time.Now()
	order.Status = domain.OrderCreated
	p.orders[order.OrderId] = order

	return nil
}

func (o *OrderRepository) FindById(orderId string) (*domain.Order, error) {
	if order, ok := o.orders[orderId]; ok {
		return order, nil
	}
	return nil, fmt.Errorf("order with id %s not found", orderId)
}

func (o *OrderRepository) Delete(order string) error {
	if _, ok := o.orders[order]; ok {
		delete(o.orders, order)
		return nil
	}
	return fmt.Errorf("order with id %s not found", order)
}

func (o *OrderRepository) GetAll() ([]*domain.Order, error) {
	var orders []*domain.Order
	for _, order := range o.orders {
		orders = append(orders, order)
	}
	return orders, nil
}

func (o *OrderRepository) UpdateStatus(orderId string, status domain.OrderStatus) error {
	if order, ok := o.orders[orderId]; ok {
		order.Status = status
	}
	return fmt.Errorf("order with id %s not found", orderId)
}
