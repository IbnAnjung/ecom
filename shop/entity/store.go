package entity

import "time"

type Store struct {
	ID           int64
	SellerUserID int64
	Name         string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}
