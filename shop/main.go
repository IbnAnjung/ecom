package main

import (
	"context"
	"edot/ecommerce/shop/app/cron"
	"edot/ecommerce/shop/app/http"
	"os"
	"os/signal"
)

func main() {

	mode := "http"
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}

	if mode == "http" {
		startHttp()
	} else if mode == "cron" {
		startCron()
	} else {
		panic("invalid args")
	}

}

func startHttp() {
	http := http.NewEchoHttpServer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		http.Start(ctx)
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()

	http.Stop()
}

func startCron() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	s := cron.NewCronServer()

	s.Start(ctx)

	<-ctx.Done()

	s.Stop()
}
