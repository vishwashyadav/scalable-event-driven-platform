package domain

import "time"

type Order struct {
	OrderId  string
	Amount   float64
	CreateAt time.Time
	Status   OrderStatus
}

type OrderStatus string

const (
	OrderCreated    OrderStatus = "OrderCreated"
	OrderProcessing OrderStatus = "OrderProcessing"
	OrderCompleted  OrderStatus = "OrderCompleted"
	OrderCancelled  OrderStatus = "OrderCancelled"
)
