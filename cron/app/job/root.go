package job

import (
	"edot/ecommerce/cron/entity"

	"github.com/go-co-op/gocron/v2"
)

func LoadJob(s gocron.Scheduler, uc entity.IOrder) {
	LoadOrderJob(s, uc)
}
