package presenter

type GetProductStockRequest struct {
	ProductIDs []string `json:"product_ids"`
}

type GetProductStockData struct {
	ProductID  string  `json:"product_id"`
	TotalStock float64 `json:"total_stock"`
}

type TransferStockRequest struct {
	SenderWarehouseID   int64                         `json:"sender_warehouse_id"`
	ReceiverWarehouseID int64                         `json:"receiver_warehouse_id"`
	Products            []TransferStockProductRequest `json:"products"`
}

type TransferStockProductRequest struct {
	ProductID string  `json:"product_id"`
	Quantity  float64 `json:"quantity"`
}
