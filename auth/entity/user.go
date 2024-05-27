package entity

import "time"

type User struct {
	ID          int64
	Name        string
	Password    string
	PhoneNumber string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type UserToken struct {
	AccessToken  string
	RefreshToken string
}
