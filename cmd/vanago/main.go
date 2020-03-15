// Vanago is a go bases web server to help create vanity imports for your custom Go code.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.lafronz.com/vanago/tools/logger"
	"go.lafronz.com/vanago/web"
)

// main is the start location for the service.
func main() {

	var wait time.Duration

	ctx := context.Background()

	svr := web.S.SingleHostProjectSetup()

	go func() {
		// start the web server on port and accept requests
		logger.Info("Server listening on port %s", web.S.Port)
		if err := svr.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, os.Interrupt)

	// Block until we receive our signal.
	<-gracefulStop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	svr.Shutdown(ctx)

	logger.Info("Server Shut Down")

	close(gracefulStop)

	os.Exit(0)
}
