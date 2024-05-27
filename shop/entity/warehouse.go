package entity

import "time"

type Warehouse struct {
	ID        int64
	StoreID   int64
	Name      string
	Status    WarehouseStatus
	CreatedAt time.Time
	UpdateAt  *time.Time
}

type WarehouseStatus int8

const (
	WarehouseStatusInActive WarehouseStatus = 0
	WarehouseStatusActive   WarehouseStatus = 1
)
