package job

import (
	"context"
	"edot/ecommerce/shop/entity"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func LoadOrderJob(ctx context.Context, s gocron.Scheduler, uc entity.IOrderUsecase) {
	// add a job to the scheduler
	fmt.Println("new job cancel order")
	_, err := s.NewJob(gocron.DurationJob(30*time.Minute), gocron.NewTask(uc.CancelExpiredOrder, ctx), gocron.WithSingletonMode(1))
	if err != nil {
		fmt.Println("err", err.Error())
	}

}
