package dto

type GetProductStockInput struct {
	RequestID  string
	ProductIDs []string
}

type ProductStock struct {
	ProductID  string
	TotalStock float64
}

type TransferStockInput struct {
	RequestID           string
	SellerUserID        int64
	SenderWarehouseID   int64
	ReceiverWarehouseID int64
	Products            []TransferStockProductInput
}

type TransferStockProductInput struct {
	ProductID string
	Quantity  float64
}
