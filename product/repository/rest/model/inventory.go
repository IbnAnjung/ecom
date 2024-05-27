package model

type GetProductStockRequest struct {
	ProductIDs []string `json:"product_ids"`
}

type GetProductStockResponseData struct {
	ProductID  string  `json:"product_id"`
	TotalStock float64 `json:"total_stock"`
}
