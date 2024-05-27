package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pkghttp "edot/ecommerce/http"
	"edot/ecommerce/nosql"
	"edot/ecommerce/product/app/http/config"
	"edot/ecommerce/product/app/http/router"
	"edot/ecommerce/product/repository/mongorepo"
	"edot/ecommerce/product/repository/rest"
	"edot/ecommerce/product/usecase/product"
	"edot/ecommerce/structvalidator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type echoHttpServer struct {
	e *echo.Echo
}

func NewEchoHttpServer() *echoHttpServer {
	return &echoHttpServer{}
}

func (server *echoHttpServer) Start(ctx context.Context) {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	config := config.LoadConfig()

	mongo, err := nosql.NewMongoConnection(ctx, nosql.MongoConfig{
		Host:     config.Mongo.Host,
		User:     config.Mongo.User,
		Password: config.Mongo.Password,
		Source:   config.Mongo.Source,
	})
	if err != nil {
		panic(fmt.Sprintf("fail connnect mongo connection: %s", err.Error()))
	}

	validator := structvalidator.NewStructValidator()

	productRepo := mongorepo.NewProductRepository(mongo.Client, mongo.Client.Database(config.Mongo.Source).Collection(config.Mongo.ProductCollection))
	inventoryRepo := rest.NewInventoryRepository(config.ServiceURI.Store)

	productUc := product.NewUsecase(productRepo, inventoryRepo, validator)
	// default http middleware
	pkghttp.LoadEchoRequiredMiddleware(e)

	router.SetupRouter(
		e, config, productUc,
	)

	server.e = e

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
}
