package model

import (
	"edot/ecommerce/shop/entity"
)

type MOrderDetail struct {
	ID         int64   `gorm:"column:id;primaryKey;auto_increment"`
	OrderID    int64   `gorm:"column:order_id"`
	ProductID  string  `gorm:"column:product_id"`
	Quantity   float64 `gorm:"column:quantity"`
	Price      float64 `gorm:"column:price"`
	TotalPrice float64 `gorm:"column:total_price"`
}

func (m *MOrderDetail) TableName() string {
	return "order_details"
}

func (m *MOrderDetail) FillFromEntity(en entity.OrderDetail) {
	m.ID = en.ID
	m.OrderID = en.OrderID
	m.ProductID = en.ProductID
	m.Price = en.Price
	m.TotalPrice = en.TotalPrice
}

func (m *MOrderDetail) ToEntity() (en entity.OrderDetail) {
	return entity.OrderDetail{
		ID:         m.ID,
		OrderID:    m.OrderID,
		ProductID:  m.ProductID,
		Price:      m.Price,
		TotalPrice: m.TotalPrice,
	}
}
