package presenter

type ValidateTokenRequest struct {
	Token string `json:"token"`
}

type ValidateTokenResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ValidateSellerTokenResponse struct {
	ID int64 `json:"id"`
}
