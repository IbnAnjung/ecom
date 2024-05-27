package entity

import (
	"context"
)

type IWarehouseRepository interface {
	GetStoreActiveWarehouse(ctx context.Context, storeId int64) (warehouses []Warehouse, err error)
	GetSellerWarehouse(ctx context.Context, sellerUserID int64) (warehouses []Warehouse, err error)
	FindByID(ctx context.Context, id int64) (warehouse Warehouse, selleUserID int64, err error)
	Update(ctx context.Context, warehouse *Warehouse) (err error)
}
