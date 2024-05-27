package main

import (
	"context"
	"os"
	"os/signal"

	"edot/ecommerce/auth/app/http"
)

func main() {

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
