package domain

import "time"

type Order struct {
	OrderId  string
	Amount   float64
	CreateAt time.Time
	Status   OrderStatus
}

type OrderStatus int

const (
	OrderCreated OrderStatus = iota
	OrderProcessing
	OrderCompleted
	OrderCancelled
)
