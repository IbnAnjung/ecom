package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pkghttp "edot/ecommerce/http"
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/app/http/config"
	"edot/ecommerce/shop/app/http/router"
	"edot/ecommerce/shop/repository/gorm"
	"edot/ecommerce/shop/usecase/order"
	"edot/ecommerce/shop/usecase/stockies"
	"edot/ecommerce/shop/usecase/warehouse"
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
	warehouseProductRepo := gorm.NewGormWarehouseProductRepository(guow)
	warehouseRepo := gorm.NewGormWarehouseRepository(guow)
	orderRepo := gorm.NewGormOrderRepository(guow)
	orderDetailRepo := gorm.NewGormOrderDetailRepository(guow)
	// usecase
	stockiesUC := stockies.NewUsecase(guow, validator, warehouseRepo, warehouseProductRepo)
	orderUc := order.NewUsecase(guow, validator, orderRepo, orderDetailRepo, warehouseRepo, warehouseProductRepo)
	whUc := warehouse.NewUsecase(guow, validator, warehouseRepo, warehouseProductRepo)

	// default http middleware
	pkghttp.LoadEchoRequiredMiddleware(e)

	router.SetupRouter(
		e, stockiesUC, orderUc, whUc, config,
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
