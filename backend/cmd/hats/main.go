package main

import (
	"backend/internal/hats/config"
	"backend/internal/hats/hatserver"
	"backend/internal/hats/repo"
	"backend/internal/hats/repo/inmem"
	"backend/internal/hats/repo/redisrepo"
	_ "backend/pkg/logging" // init logrus
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

	hatRepo := initRepo()
	defer hatRepo.Close()

	// middleware filter chain
	hooks := twirp.ChainHooks(repo.Hook(hatRepo))

	twirpHandler := hatspb.NewHatsServer(&hatserver.Server{}, hooks)

	srv := &http.Server{
		Addr:    config.ListenAddress,
		Handler: twirpHandler,
	}

	srv.Handler = twirpHandler

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

	logrus.Infof("main() %s listening on %s", config.AppName, config.ListenAddress)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logrus.Fatalf("HTTP server ListenAndServe err=%v", err)
	}

	<-idleConnsClosed

	logrus.Infof("main() %s shut down", config.AppName)
}

func initRepo() repo.HatRepo {

	switch config.DatastoreConfig {
	case "inmem":
		return inmem.NewRepo()
	case "redis":
		return redisrepo.NewRepo()
	default:
		logrus.Panicf("unexpected config.DatastoreConfig=%s", config.DatastoreConfig)
		return nil
	}
}
