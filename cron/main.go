package main

import (
	"context"
	"edot/ecommerce/cron/app"
	"os"
	"os/signal"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	s := app.NewCronServer()

	s.Start()

	<-ctx.Done()

	s.Stop()
}
