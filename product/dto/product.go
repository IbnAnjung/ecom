package dto

import "time"

type GetListProductInput struct {
	RequestID string
	Keyword   string
	Limit     int8
	Page      int16
}

type FindProductOutput struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Stock       float64
	Price       float64
}
