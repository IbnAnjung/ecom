package repository

import (
	"context"
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/repository/gorm/model"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/orm"
	"errors"

	"gorm.io/gorm"
)

type sellerUserRepository struct {
	uow orm.IGormUow
}

func NewGormSellerUserRepository(
	uow orm.IGormUow,
) entity.ISellerUserRepository {
	return &sellerUserRepository{
		uow,
	}
}

func (r *sellerUserRepository) FindUserByUsername(ctx context.Context, username string) (u entity.SellerUser, err error) {
	m := model.MSellerUser{}
	db := r.uow.GetDB().WithContext(ctx)

	if err = db.WithContext(ctx).Where("username = ? ", username).Find(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	u = m.ToEntity()

	return
}

func (r *sellerUserRepository) Create(ctx context.Context, u *entity.SellerUser) (err error) {
	m := model.MSellerUser{}
	m.FillFromEntity(*u)

	if err = r.uow.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	u.ID = m.ID

	return nil
}

func (r *sellerUserRepository) FindById(ctx context.Context, id int64) (user entity.SellerUser, err error) {
	m := model.MSellerUser{}
	if err = r.uow.GetDB().WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "user tidak valid")
		}
		err = e
		return
	}

	user = m.ToEntity()
	return
}
