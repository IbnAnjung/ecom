package entity

type SellerUser struct {
	ID       int64
	Username string
	Password string
}

type SellerUserToken struct {
	AccessToken  string
	RefreshToken string
}
