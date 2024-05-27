package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"edot/ecommerce/auth/app/http/config"
	"edot/ecommerce/auth/app/http/router"
	"edot/ecommerce/auth/internal/jwt"
	repository "edot/ecommerce/auth/repository/gorm"
	"edot/ecommerce/auth/usecase/auth"
	"edot/ecommerce/auth/usecase/seller_user"
	"edot/ecommerce/auth/usecase/user"
	"edot/ecommerce/crypt"
	pkghttp "edot/ecommerce/http"
	"edot/ecommerce/orm"
	"edot/ecommerce/sql"
	"edot/ecommerce/structvalidator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type echoHttpServer struct {
	e     *echo.Echo
	mysql sql.MysqlConnection
}

func NewEchoHttpServer() *echoHttpServer {
	return &echoHttpServer{}
}

func (server *echoHttpServer) Start(ctx context.Context) {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	config := config.LoadConfig()
	t, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(fmt.Sprintf("server startup panic: %s", err))
	}

	validator := structvalidator.NewStructValidator()
	hasher := crypt.NewHasherString()

	jwt := jwt.NewJwtServive(jwt.JwtConfig{
		SecretKey:            config.Jwt.SecretKey,
		SellerSecretKey:      config.Jwt.SellerSecretKey,
		AccessTokenLifeTime:  time.Duration(config.Jwt.AccessTokenLifeTime) * time.Hour,
		RefreshTokenLifeTime: time.Duration(config.Jwt.RefreshTokenLifeTime) * time.Hour,
	})

	// open mysql connection
	mysql, err := sql.NewMysqlConnection(ctx, sql.MysqlConfig{
		User:               config.Mysql.User,
		Password:           config.Mysql.Password,
		Host:               config.Mysql.Host,
		Port:               config.Mysql.Port,
		DbName:             config.Mysql.Schema,
		Loc:                t,
		Timeout:            time.Duration(config.Mysql.Timeout) * time.Second,
		MaxIddleConnection: config.Mysql.MaxIddleConnection,
		MaxOpenConnection:  config.Mysql.MaxOpenConnection,
		MaxLifeTime:        config.Mysql.MaxLifeTime,
	})
	if err != nil {
		panic(fmt.Sprintf("server startup panic: %s", err))
	}

	guow, err := orm.NewGormOrm(orm.GormConfig{
		Connection: mysql.Db,
		Dialect:    orm.MysqlDialect,
	})

	if err != nil {
		panic(fmt.Sprintf("server startup panic: %s", err))
	}

	// repository
	userRepository := repository.NewGormUserRepository(guow)
	sellerUserRepository := repository.NewGormSellerUserRepository(guow)

	// usecase
	userUc := user.NewUsecase(hasher, jwt, validator, userRepository)
	sellerUserUc := seller_user.NewUsecase(hasher, jwt, validator, sellerUserRepository)
	authUc := auth.NewUsecase(userRepository, sellerUserRepository, jwt)

	// default http middleware
	pkghttp.LoadEchoRequiredMiddleware(e)

	router.SetupRouter(
		e, userUc, sellerUserUc, authUc,
	)

	server.e = e
	server.mysql = mysql

	if err := e.Start(fmt.Sprintf(":%s", config.Http.Port)); err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("server startup panic: %s", err))
	}

}

func (server *echoHttpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.e.Shutdown(ctx); err != nil {
		server.e.Logger.Fatal(err)
	}

	if err := server.mysql.Cleanup(); err != nil {
		server.e.Logger.Fatal(err)
	}
}
