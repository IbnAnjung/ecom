package model

import "edot/ecommerce/shop/dto"

type MProductStok struct {
	ID         string  `gorm:"id"`
	TotalStock float64 `gorm:"total_stock"`
}

func (m *MProductStok) ToEntity() dto.ProductStock {
	return dto.ProductStock{
		ProductID:  m.ID,
		TotalStock: m.TotalStock,
	}
}
