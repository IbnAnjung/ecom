package job

import (
	"context"
	"edot/ecommerce/shop/entity"

	"github.com/go-co-op/gocron/v2"
)

func LoadJob(ctx context.Context, s gocron.Scheduler, uc entity.IOrderUsecase) {
	LoadOrderJob(ctx, s, uc)
}
