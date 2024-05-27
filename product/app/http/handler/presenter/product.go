package presenter

type GetListProductRequest struct {
	KeyWord string `query:"keyword"`
	Page    int16  `query:"page"`
	Limit   int8   `query:"limit"`
}

type GetListProductRespnose struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       float64 `json:"stock"`
	Price       float64 `json:"price"`
}
