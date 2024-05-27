package gorm

import (
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/entity"
	"edot/ecommerce/shop/repository/gorm/model"
	"errors"

	"github.com/jinzhu/gorm"
)

type warehouseRepository struct {
	uow orm.IGormUow
}

func NewGormWarehouseRepository(
	uow orm.IGormUow,
) entity.IWarehouseRepository {
	return &warehouseRepository{
		uow,
	}
}

func (r *warehouseRepository) GetStoreActiveWarehouse(ctx context.Context, storeId int64) (warehouses []entity.Warehouse, err error) {
	m := []model.MWarehouse{}

	if err = r.uow.GetDB().WithContext(ctx).
		Where("store_id = ?", storeId).
		Where("status = ?", entity.WarehouseStatusActive).
		Find(&m).Error; err != nil {
		return
	}

	warehouses = make([]entity.Warehouse, len(m))
	for i, v := range m {
		warehouses[i] = v.ToEntity()
	}

	return
}

func (r *warehouseRepository) GetSellerWarehouse(ctx context.Context, sellerUserID int64) (warehouses []entity.Warehouse, err error) {
	m := []model.MWarehouse{}

	if err = r.uow.GetDB().WithContext(ctx).
		Select("w.*").
		Table("stores s").
		Joins("JOIN warehouses w ON w.store_id = s.id").
		Where("s.seller_user_id = ?", sellerUserID).
		Find(&m).Error; err != nil {
		return
	}

	warehouses = make([]entity.Warehouse, len(m))
	for i, v := range m {
		warehouses[i] = v.ToEntity()
	}

	return
}

func (r *warehouseRepository) FindByID(ctx context.Context, id int64) (warehouse entity.Warehouse, selleUserID int64, err error) {
	m := model.MWarehouseWithUserId{}

	if err = r.uow.GetDB().WithContext(ctx).
		Select("warehouses.*, s.seller_user_id").
		Joins("JOIN stores s ON warehouses.store_id = s.id").
		Where("warehouses.id = ?", id).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "warehouse not found")
		}
		return
	}

	warehouse = m.ToEntity()
	selleUserID = m.SellerUserID

	return
}

func (r *warehouseRepository) Update(ctx context.Context, warehouse *entity.Warehouse) (err error) {
	if warehouse == nil {
		return
	}

	m := model.MWarehouse{}
	m.FillFromEntity(*warehouse)

	if err = r.uow.GetDB().WithContext(ctx).Updates(&m).Error; err != nil {
		return
	}

	return
}
