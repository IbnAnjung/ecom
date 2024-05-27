package job

import (
	"edot/ecommerce/cron/entity"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func LoadOrderJob(s gocron.Scheduler, uc entity.IOrder) {
	// add a job to the scheduler
	_, err := s.NewJob(gocron.DurationJob(1*time.Second), gocron.NewTask(uc.CancelExpiredOrder), gocron.WithSingletonMode(1))
	if err != nil {
		fmt.Println("err", err.Error())
	}

}
