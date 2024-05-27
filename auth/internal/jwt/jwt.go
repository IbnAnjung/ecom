package jwt

import (
	coreerror "edot/ecommerce/error"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IJwtService interface {
	GenerateAccessToken(claim UserClaim, userType UserType) (token string, err error)
	GenerateRefreshToken(claim UserClaim, userType UserType) (token string, err error)
	ValidateToken(tokenString string, userType UserType) (c *UserClaim, err error)
}

type UserType int8

const (
	UserTypeUser   UserType = 0
	UserTypeSeller UserType = 1
)

type UserClaim struct {
	UserID   int64  `json:"id"`
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

type JwtConfig struct {
	SecretKey            string
	SellerSecretKey      string
	AccessTokenLifeTime  time.Duration
	RefreshTokenLifeTime time.Duration
}

type jwtService struct {
	conf JwtConfig
}

func NewJwtServive(conf JwtConfig) IJwtService {
	return &jwtService{
		conf,
	}
}

func (s *jwtService) getSecretKey(userType UserType) string {
	if userType == UserTypeSeller {
		return s.conf.SellerSecretKey
	} else {
		return s.conf.SecretKey
	}
}

func (s *jwtService) GenerateAccessToken(claim UserClaim, userType UserType) (token string, err error) {
	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(s.conf.AccessTokenLifeTime))
	claim.IssuedAt = jwt.NewNumericDate(time.Now())

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = jwtToken.SignedString([]byte(s.getSecretKey(userType)))
	if err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	return
}

func (s *jwtService) GenerateRefreshToken(claim UserClaim, userType UserType) (token string, err error) {
	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(s.conf.RefreshTokenLifeTime))
	claim.IssuedAt = jwt.NewNumericDate(time.Now())

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = jwtToken.SignedString([]byte(s.getSecretKey(userType)))
	if err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	return
}

func (s *jwtService) ValidateToken(tokenString string, userType UserType) (c *UserClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.getSecretKey(userType)), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, err.Error())
		err = e
		return
	} else if claim, ok := token.Claims.(*UserClaim); ok {
		return claim, nil
	} else {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "unknown error")
		err = e
		return
	}
}
