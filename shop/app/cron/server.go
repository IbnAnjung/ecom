package cron

import (
	"context"
	"edot/ecommerce/orm"
	"edot/ecommerce/shop/app/cron/job"
	"edot/ecommerce/shop/app/http/config"
	"edot/ecommerce/shop/repository/gorm"
	"edot/ecommerce/shop/usecase/order"
	"edot/ecommerce/sql"
	"edot/ecommerce/structvalidator"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type cronServer struct {
	s     gocron.Scheduler
	mysql sql.MysqlConnection
}

func NewCronServer() *cronServer {
	return &cronServer{}
}

func (c *cronServer) Start(ctx context.Context) {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic("error start scheduler")
	}

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
	orderUc := order.NewUsecase(guow, validator, orderRepo, orderDetailRepo, warehouseRepo, warehouseProductRepo)

	job.LoadJob(ctx, s, orderUc)

	c.s = s
	s.Start()
}

func (c *cronServer) Stop() {
	c.s.Shutdown()
}
