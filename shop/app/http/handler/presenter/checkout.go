package presenter

type CheckoutRequest struct {
	StoreID  int64                    `json:"store_id"`
	Products []CheckoutProductRequest `json:"products"`
}

type CheckoutProductRequest struct {
	ID       string  `json:"id"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
}
