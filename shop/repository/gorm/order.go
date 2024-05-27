package gorm

import (
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/entity"
	"edot/ecommerce/shop/repository/gorm/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type orderRepository struct {
	uow orm.IGormUow
}

func NewGormOrderRepository(
	uow orm.IGormUow,
) entity.IOrderRepository {
	return &orderRepository{
		uow,
	}
}

func (r *orderRepository) CreateOrder(ctx context.Context, input *entity.Order) (err error) {
	m := model.MOrder{}
	m.FillFromEntity(*input)

	if err = r.uow.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		return
	}

	input.ID = m.ID
	return
}

func (r *orderRepository) FindOrder(ctx context.Context, id int64) (order entity.Order, err error) {
	m := model.MOrder{}

	if err = r.uow.GetDB().WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "")
		}
		return
	}

	order = m.ToEntity()

	return
}

func (r *orderRepository) GetExpiredOrder(ctx context.Context) (orders []entity.Order, err error) {
	m := []model.MOrder{}

	if err = r.uow.GetDB().WithContext(ctx).Where("status = ?", entity.OrderStatusCreated).
		Where("expired_time <= ?", time.Now().Format("2006-01-02 15:04:05")).
		Find(&m).Order("id").Limit(50).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "")
		}
		return
	}

	orders = make([]entity.Order, len(m))
	for i, v := range m {
		orders[i] = v.ToEntity()
	}

	return
}

func (r *orderRepository) Update(ctx context.Context, order *entity.Order) (err error) {
	m := model.MOrder{}
	m.FillFromEntity(*order)
	if err = r.uow.GetDB().WithContext(ctx).Updates(&m).Error; err != nil {
		return
	}

	return
}
