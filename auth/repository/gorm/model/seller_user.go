package model

import (
	"edot/ecommerce/auth/entity"
)

type MSellerUser struct {
	ID       int64  `gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (m *MSellerUser) TableName() string {
	return "seller_users"
}

func (m *MSellerUser) ToEntity() (en entity.SellerUser) {
	en.ID = m.ID
	en.Username = m.Username
	en.Password = m.Password
	return
}

func (m *MSellerUser) FillFromEntity(en entity.SellerUser) {
	m.ID = en.ID
	m.Username = en.Username
	m.Password = en.Password
}
