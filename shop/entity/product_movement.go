package entity

import "time"

type ProductMovement struct {
	ID                 int64
	WarehouseProductID int64
	Reference          string
	ReferenceNumber    ProductMovementReference
	ReferenceTime      time.Time
	Quantity           float64
	CreatedAt          time.Time
}

type ProductMovementReference int8

const (
	ProductMovementReferenceOrder         = 1
	ProductMovementReferenceTransferStock = 2
)
