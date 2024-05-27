package model

import "edot/ecommerce/shop/entity"

type MWarehouse struct {
	ID      int64  `gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	StoreID int64  `gorm:"column:store_id"`
	Name    string `gorm:"column:name"`
	Status  *int8  `gorm:"column:status"`
}

func (m *MWarehouse) TableName() string {
	return "warehouses"
}

func (m *MWarehouse) ToEntity() entity.Warehouse {
	return entity.Warehouse{
		ID:      m.ID,
		StoreID: m.StoreID,
		Name:    m.Name,
		Status:  entity.WarehouseStatus(*m.Status),
	}
}

func (m *MWarehouse) FillFromEntity(en entity.Warehouse) {
	m.ID = en.ID
	m.StoreID = en.StoreID
	m.Name = en.Name
	m.Status = (*int8)(&en.Status)
}

type MWarehouseWithUserId struct {
	MWarehouse
	SellerUserID int64 `gorm:"column:seller_user_id"`
}
