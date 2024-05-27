package model

import (
	"edot/ecommerce/shop/entity"
	"time"
)

type MOrder struct {
	ID          int64      `gorm:"column:id;primaryKey;auto_increment"`
	StoreID     int64      `gorm:"column:store_id"`
	UserID      int64      `column:"user_id"`
	TotalPrice  float64    `column:"total_price"`
	CreatedTime time.Time  `column:"created_time"`
	ExpiredTime time.Time  `column:"expired_time"`
	PaymentTime *time.Time `column:"payment_time"`
	Status      int8       `column:"status"`
}

func (m *MOrder) TableName() string {
	return "orders"
}

func (m *MOrder) FillFromEntity(en entity.Order) {
	m.ID = en.ID
	m.StoreID = en.StoreID
	m.UserID = en.UserID
	m.TotalPrice = en.TotalPrice
	m.CreatedTime = en.CreatedTime
	m.ExpiredTime = en.ExpiredTime
	m.PaymentTime = en.PaymentTime
	m.Status = int8(en.Status)
}

func (m *MOrder) ToEntity() (en entity.Order) {
	return entity.Order{
		ID:          m.ID,
		StoreID:     m.StoreID,
		UserID:      m.UserID,
		TotalPrice:  m.TotalPrice,
		CreatedTime: m.CreatedTime,
		ExpiredTime: m.ExpiredTime,
		PaymentTime: m.PaymentTime,
		Status:      entity.OrderStatus(m.Status),
	}
}
