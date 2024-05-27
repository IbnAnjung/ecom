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

type userRepository struct {
	uow orm.IGormUow
}

func NewGormUserRepository(
	uow orm.IGormUow,
) entity.IUserRepository {
	return &userRepository{
		uow,
	}
}

func (r *userRepository) FindUserByEmailOrPhoneNumber(ctx context.Context, phoneNumber, email *string) (u entity.User, err error) {
	if phoneNumber == nil && email == nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "empty phone_number and email")
		err = e
		return
	}

	m := model.MUser{}
	db := r.uow.GetDB().WithContext(ctx)
	if phoneNumber != nil && email != nil {
		db = db.Where("phone_number = ?", phoneNumber).
			Or("email = ?", email)
	} else if phoneNumber != nil {
		db = db.Where("phone_number = ?", phoneNumber)
	} else if email != nil {
		db = db.Where("email = ?", email)
	}

	if err = db.WithContext(ctx).Find(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	u = m.ToEntity()

	return
}

func (r *userRepository) Create(ctx context.Context, u *entity.User) (err error) {
	m := model.MUser{}
	m.FillFromEntity(*u)

	if err = r.uow.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	u.ID = m.ID

	return nil
}

func (r *userRepository) FindById(ctx context.Context, id int64) (user entity.User, err error) {
	m := model.MUser{}
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
