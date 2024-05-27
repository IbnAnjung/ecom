package stockies

import (
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/entity"
	"edot/ecommerce/structvalidator"
)

type stockies struct {
	gormUow              orm.IGormUow
	validator            structvalidator.IStructValidator
	warehouseRepo        entity.IWarehouseRepository
	warehouseProductRepo entity.IWarehouseProductRepository
}

func NewUsecase(
	gormUow orm.IGormUow,
	validator structvalidator.IStructValidator,
	warehouseRepo entity.IWarehouseRepository,
	warehouseProductRepo entity.IWarehouseProductRepository,
) entity.IStockiesUsecase {
	return &stockies{
		gormUow,
		validator,
		warehouseRepo,
		warehouseProductRepo,
	}
}
