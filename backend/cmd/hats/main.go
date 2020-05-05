package main

import (
	"backend/internal/hats/config"
	"backend/internal/hats/hatrabbit"
	"backend/internal/hats/hatserver"
	"backend/internal/hats/mongoclient"
	"backend/pkg/httpwrap"
	_ "backend/pkg/logging" // init logrus
	"backend/pkg/rabbit"
	"backend/rpc/hatspb"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/twitchtv/twirp"

	"github.com/sirupsen/logrus"
)

// hats
func main() {

	logrus.Infof("main() %s starting", config.AppName)

	// middleware filter chain
	hooks := twirp.ChainHooks(mongoclient.ServerHooks(), rabbit.ServerHooks())

	twirpHandler := hatspb.NewHatsServer(hatserver.NewServer(), hooks)

	// grabs the http headers
	wrapped := httpwrap.WithHeaders(twirpHandler)

	srv := &http.Server{
		Addr:    config.ListenAddress,
		Handler: wrapped,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logrus.Infof("HTTP server Shutdown err=%v", err)
		}
		close(idleConnsClosed)
	}()

	// start rabbit listeners
	hatrabbit.Listen()

	logrus.Infof("main() %s listening on %s", config.AppName, config.ListenAddress)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logrus.Fatalf("HTTP server ListenAndServe err=%v", err)
	}

	<-idleConnsClosed

	logrus.Infof("main() %s shut down", config.AppName)
}
