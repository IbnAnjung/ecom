package gorm

import (
	"context"
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/entity"
	"edot/ecommerce/shop/repository/gorm/model"
)

type orderDetailRepository struct {
	uow orm.IGormUow
}

func NewGormOrderDetailRepository(
	uow orm.IGormUow,
) entity.IOrderDetailRepository {
	return &orderDetailRepository{
		uow,
	}
}

func (r *orderDetailRepository) CreateBulk(ctx context.Context, input *[]entity.OrderDetail) (err error) {
	m := make([]model.MOrderDetail, len(*input))
	for i, v := range *input {
		m[i] = model.MOrderDetail{}
		m[i].FillFromEntity(v)
	}

	if err = r.uow.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		return
	}

	en := make([]entity.OrderDetail, len(m))
	for i, v := range m {
		en[i] = v.ToEntity()
	}

	input = &en

	return
}

func (r *orderDetailRepository) GetDetails(ctx context.Context, orderID int64) (details []entity.OrderDetail, err error) {
	m := []model.MOrderDetail{}

	if err = r.uow.GetDB().WithContext(ctx).Where("order_id = ?", orderID).Find(&m).Error; err != nil {
		return
	}

	details = make([]entity.OrderDetail, len(m))
	for i, v := range m {
		details[i] = v.ToEntity()
	}

	return
}
