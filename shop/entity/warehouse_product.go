package entity

type WarehouseProduct struct {
	ID          int64
	WarehouseID int64
	ProductID   string
	Stock       float64
}
