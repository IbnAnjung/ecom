package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

type IHash interface {
	HashString(plainText string) (hText string, err error)
	CompareHash(hText, plainText string) error
}

type hashString struct{}

func NewHasherString() IHash {
	return &hashString{}
}

func (h *hashString) HashString(plainText string) (hText string, err error) {
	hByte, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hText = string(hByte)
	return
}

func (h *hashString) CompareHash(hText, plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hText), []byte(plainText))
}
