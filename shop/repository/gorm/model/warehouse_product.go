package model

import "edot/ecommerce/shop/entity"

type MWarehouseProduct struct {
	ID          int64   `gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	WarehouseID int64   `gorm:"column:warehouse_id"`
	ProductID   string  `gorm:"column:product_id"`
	Stock       float64 `gorm:"column:stock"`
}

func (m *MWarehouseProduct) TableName() string {
	return "warehouse_products"
}

func (m *MWarehouseProduct) ToEntity() entity.WarehouseProduct {
	return entity.WarehouseProduct{
		ID:          m.ID,
		WarehouseID: m.WarehouseID,
		ProductID:   m.ProductID,
		Stock:       m.Stock,
	}
}

func (m *MWarehouseProduct) FillFromEntity(en entity.WarehouseProduct) {
	m.ID = en.ID
	m.WarehouseID = en.WarehouseID
	m.ProductID = en.ProductID
	m.Stock = en.Stock
}
