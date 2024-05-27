package product

import (
	"context"
	"edot/ecommerce/product/dto"
	"edot/ecommerce/product/entity"
	"fmt"
)

type GetListProductInput struct {
	RequestID string `validate:"required"`
	Keyword   string `validate:"omitempty,max=50"`
	Limit     int8   `validate:"omitempty,numeric"`
	Page      int16  `validate:"omitempty,numeric"`
}

func (uc *productUsecase) GetListProduct(ctx context.Context, input dto.GetListProductInput) (products []dto.FindProductOutput, err error) {
	if err = uc.validator.Validate(GetListProductInput{
		RequestID: input.RequestID,
		Keyword:   input.Keyword,
		Limit:     input.Limit,
		Page:      input.Page,
	}); err != nil {
		fmt.Printf("erro validation data, request_id: %s", input.RequestID)
		return
	}

	if input.Limit == 0 {
		input.Limit = 10
	}

	if input.Page == 0 {
		input.Page = 1
	}

	// find product
	enProducts, err := uc.productRepository.Find(ctx, input)
	if err != nil {
		fmt.Printf("fail get list product: %s %s", input.RequestID, err.Error())
		return
	}

	// get inventories
	productIds := make([]string, len(enProducts))
	for i, v := range enProducts {
		productIds[i] = v.ID
	}
	mapProductRepositories := map[string]entity.ProductInventory{}
	productInventories, err := uc.inventoryRepository.GetProductInventories(ctx, productIds)
	if err != nil {
		fmt.Printf("fail get list product: %s %s", input.RequestID, err.Error())
		return
	}

	for _, v := range productInventories {
		mapProductRepositories[v.ID] = v
	}

	for _, v := range enProducts {
		temp := dto.FindProductOutput{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		}

		if inv, ok := mapProductRepositories[v.ID]; ok {
			temp.Stock = inv.Stock
			temp.Price = 0
		}

		products = append(products, temp)
	}

	return
}
