package gorm

import (
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/dto"
	"edot/ecommerce/shop/entity"
	"edot/ecommerce/shop/repository/gorm/model"

	"gorm.io/gorm/clause"
)

type warehouseProductRepository struct {
	uow orm.IGormUow
}

func NewGormWarehouseProductRepository(
	uow orm.IGormUow,
) entity.IWarehouseProductRepository {
	return &warehouseProductRepository{
		uow,
	}
}

func (r *warehouseProductRepository) GetProductTotalStock(ctx context.Context, productIDs []string) (stocks []dto.ProductStock, err error) {
	m := []model.MProductStok{}

	if err = r.uow.GetDB().WithContext(ctx).Table("warehouse_products").
		Select("product_id id, sum(stock) total_stock").
		Where("product_id IN (?)", productIDs).
		Group("product_id").
		Find(&m).Error; err != nil {
		return
	}

	stocks = make([]dto.ProductStock, len(m))
	for i, v := range m {
		stocks[i] = v.ToEntity()
	}

	return
}

func (r *warehouseProductRepository) GetForUpdateProccess(ctx context.Context, warehouseId []int64, productIDs []string) (whProducts []entity.WarehouseProduct, err error) {
	if len(warehouseId) == 0 && len(productIDs) == 0 {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "not allowed without condition")
		return
	}

	m := []model.MWarehouseProduct{}
	db := r.uow.GetDB().WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"})
	if len(warehouseId) > 0 {
		db = db.Where("warehouse_id in (?)", warehouseId)
	}
	if len(productIDs) > 0 {
		db = db.Where("product_id in (?)", productIDs)
	}
	if err = db.Find(&m).Error; err != nil {
		return
	}

	whProducts = make([]entity.WarehouseProduct, len(m))
	for i, v := range m {
		whProducts[i] = v.ToEntity()
	}

	return
}

func (r *warehouseProductRepository) UpdateStock(ctx context.Context, whProducts []entity.WarehouseProduct) (err error) {
	data := make([]model.MWarehouseProduct, len(whProducts))
	for i, v := range whProducts {
		m := model.MWarehouseProduct{}
		m.FillFromEntity(v)
		data[i] = m
	}

	if err = r.uow.GetDB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"stock"}),
	}).Create(&data).Error; err != nil {
		return
	}
	return
}
