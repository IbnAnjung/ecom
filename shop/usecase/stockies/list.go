package stockies

import (
	"context"
	"edot/ecommerce/shop/dto"
	"fmt"
)

type GetProductStock struct {
	RequestID  string   `validate:"required"`
	ProductIDs []string `validate:"required,dive"`
}

func (uc *stockies) GetProductStok(ctx context.Context, input dto.GetProductStockInput) (stocks []dto.ProductStock, err error) {
	if err = uc.validator.Validate(GetProductStock{
		RequestID:  input.RequestID,
		ProductIDs: input.ProductIDs,
	}); err != nil {
		fmt.Printf("validation fail, %s %v\n", input.RequestID, err)
		return
	}

	stocks, err = uc.warehouseProductRepo.GetProductTotalStock(ctx, input.ProductIDs)
	if err != nil {
		fmt.Printf("fail get warehouse product, %s %v\n", input.RequestID, err)
		return
	}

	return
}
