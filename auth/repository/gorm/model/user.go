package model

import (
	"edot/ecommerce/auth/entity"
	"time"
)

type MUser struct {
	ID          int64      `gorm:"column:id;primaryKey;AUTO_INCREMENT"`
	Name        string     `gorm:"column:name"`
	PhoneNumber string     `gorm:"column:phone_number"`
	Email       string     `gorm:"column:email"`
	Password    string     `gorm:"column:password"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (m *MUser) TableName() string {
	return "users"
}

func (m *MUser) ToEntity() (en entity.User) {
	en.ID = m.ID
	en.Name = m.Name
	en.PhoneNumber = m.PhoneNumber
	en.Password = m.Password
	en.Email = m.Email
	en.CreatedAt = m.CreatedAt
	en.UpdatedAt = m.UpdatedAt
	return
}

func (m *MUser) FillFromEntity(en entity.User) {
	m.ID = en.ID
	m.Name = en.Name
	m.PhoneNumber = en.PhoneNumber
	m.Password = en.Password
	m.Email = en.Email
	m.CreatedAt = en.CreatedAt
	m.UpdatedAt = en.UpdatedAt
}
