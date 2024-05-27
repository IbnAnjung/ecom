package entity

import "time"

type Order struct {
	ID          int64
	StoreID     int64
	UserID      int64
	TotalPrice  float64
	CreatedTime time.Time
	ExpiredTime time.Time
	PaymentTime *time.Time
	Status      OrderStatus
}

type OrderStatus int8

const (
	OrderStatusCreated OrderStatus = 0
	OrderStatusPaid    OrderStatus = 1
	OrderStatusExpired OrderStatus = 9
)
