package model

import (
	"edot/ecommerce/product/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MProduct struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	StoreID     int64              `bson:"store_id"`
}

func (m *MProduct) ToEntity() entity.Product {
	return entity.Product{
		ID:          m.ID.Hex(),
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		StoreID:     m.StoreID,
	}
}
