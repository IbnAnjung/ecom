package warehouse

import (
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/shop/dto"
	"edot/ecommerce/shop/entity"
	"fmt"
)

type ToogleStatusWarehouse struct {
	RequestID    string `validate:"required"`
	SellerUserID int64  `validate:"required,number"`
	WarehouseID  int64  `validate:"required,number"`
}

func (uc *warehouseUsecase) ToggleStatus(ctx context.Context, input dto.ToogleStatusWarehouseInput) (err error) {
	if err = uc.validator.Validate(ToogleStatusWarehouse{
		RequestID:    input.RequestID,
		SellerUserID: input.SellerUserID,
		WarehouseID:  input.WarehouseID,
	}); err != nil {
		fmt.Printf("error validate input %s %s", input.RequestID, err.Error())
		return
	}

	warehouse, sellerUserID, err := uc.warehouseRepository.FindByID(ctx, input.WarehouseID)
	if err != nil {
		return
	}

	fmt.Println(input.SellerUserID, sellerUserID)
	if input.SellerUserID != sellerUserID {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeForbidden, "You dont have any access")
		return
	}

	if warehouse.Status == entity.WarehouseStatusActive {
		warehouse.Status = entity.WarehouseStatusInActive
	} else {
		warehouse.Status = entity.WarehouseStatusActive
	}

	if err = uc.warehouseRepository.Update(ctx, &warehouse); err != nil {
		return
	}

	return
}
