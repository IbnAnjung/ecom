package warehouse

import (
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/entity"
	"edot/ecommerce/structvalidator"
)

type warehouseUsecase struct {
	gormUow                    orm.IGormUow
	validator                  structvalidator.IStructValidator
	warehouseRepository        entity.IWarehouseRepository
	warehouseProductRepository entity.IWarehouseProductRepository
}

func NewUsecase(
	gormUow orm.IGormUow,
	validator structvalidator.IStructValidator,
	warehouseRepository entity.IWarehouseRepository,
	warehouseProductRepository entity.IWarehouseProductRepository,
) entity.IWarehouseUsecase {
	return &warehouseUsecase{
		gormUow,
		validator,
		warehouseRepository,
		warehouseProductRepository,
	}
}
