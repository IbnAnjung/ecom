package order

import (
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/entity"
	"edot/ecommerce/structvalidator"
)

type orderUsecase struct {
	gormUow                    orm.IGormUow
	validator                  structvalidator.IStructValidator
	orderRepository            entity.IOrderRepository
	orderDetailRepository      entity.IOrderDetailRepository
	warehouseRepository        entity.IWarehouseRepository
	warehouseProductRepository entity.IWarehouseProductRepository
}

func NewUsecase(
	gormUow orm.IGormUow,
	validator structvalidator.IStructValidator,
	orderRepository entity.IOrderRepository,
	orderDetailRepository entity.IOrderDetailRepository,
	warehouseRepository entity.IWarehouseRepository,
	warehouseProductRepository entity.IWarehouseProductRepository,
) entity.IOrderUsecase {
	return &orderUsecase{
		gormUow,
		validator,
		orderRepository,
		orderDetailRepository,
		warehouseRepository,
		warehouseProductRepository,
	}
}
