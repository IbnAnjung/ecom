package product

import (
	"edot/ecommerce/product/entity"
	"edot/ecommerce/structvalidator"
)

type productUsecase struct {
	productRepository   entity.IProductRepository
	inventoryRepository entity.IInventoryRepository
	validator           structvalidator.IStructValidator
}

func NewUsecase(
	productRepository entity.IProductRepository,
	inventoryRepository entity.IInventoryRepository,
	validator structvalidator.IStructValidator,
) entity.IProductUsecase {
	return &productUsecase{
		productRepository, inventoryRepository, validator,
	}
}
