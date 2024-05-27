package app

import (
	"edot/ecommerce/cron/app/job"
	"edot/ecommerce/cron/usecase/order"
	"edot/ecommerce/sql"

	"github.com/go-co-op/gocron/v2"
)

type cronServer struct {
	s     gocron.Scheduler
	mysql sql.MysqlConnection
}

func NewCronServer() *cronServer {
	return &cronServer{}
}

func (c *cronServer) Start() {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic("error start scheduler")
	}

	orderUc := order.NewUsecase()

	job.LoadJob(s, orderUc)

	c.s = s
	s.Start()
}

func (c *cronServer) Stop() {
	c.s.Shutdown()
}
