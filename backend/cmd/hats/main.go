package main

import (
	"backend/internal/hats/config"
	"backend/internal/hats/server"
	_ "backend/pkg/logging" // init logrus
	"backend/rpc/hatspb"
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
)

// Run the implementation in a local server
func main() {

	twirpHandler := hatspb.NewHatsServer(&server.Server{}, nil)

	srv := &http.Server{
		Addr:    config.ListenAddress,
		Handler: twirpHandler,
	}

	srv.Handler = twirpHandler

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logrus.Infof("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	logrus.Infof("main() %s listening on %s", config.AppName, config.ListenAddress)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logrus.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed

}
