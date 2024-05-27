package rest

import (
	"bytes"
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/product/entity"
	"edot/ecommerce/product/repository/rest/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type inventory struct {
	baseURI string
}

func NewInventoryRepository(baseURI string) entity.IInventoryRepository {
	return &inventory{baseURI}
}

func (r *inventory) GetProductInventories(ctx context.Context, ids []string) (inv []entity.ProductInventory, err error) {
	jsonBody, _ := json.Marshal(model.GetProductStockRequest{
		ProductIDs: ids,
	})
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/stock", r.baseURI), bytes.NewReader(jsonBody))
	if err != nil {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		return
	}

	resBody := model.StandartSuccessResponse{}
	data := []model.GetProductStockResponseData{}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		return
	}

	if err = json.Unmarshal(bodyBytes, &resBody); err != nil {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		return
	}

	dataByte, err := json.Marshal(resBody.Data)
	if err != nil {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		return
	}

	if err = json.Unmarshal(dataByte, &data); err != nil {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		return
	}

	inv = make([]entity.ProductInventory, len(data))
	for i, v := range data {
		inv[i] = entity.ProductInventory{
			ID:    v.ProductID,
			Stock: v.TotalStock,
		}
	}

	return
}
