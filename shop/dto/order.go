package dto

type CheckoutOrderInput struct {
	RequestID string
	StoreID   int64
	UserID    int64
	Products  []CheckoutOrderProductInput
}

type CheckoutOrderProductInput struct {
	ProductID string
	Quantity  float64
	Price     float64
}

type CheckoutInput struct {
	ProductID string
	Quantity  float64
}
