package presenter

type SellerUserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SellerUserRegisterResponse struct {
	ID           int64  `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SellerUserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SellerUserLoginResponse struct {
	ID           int64  `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
