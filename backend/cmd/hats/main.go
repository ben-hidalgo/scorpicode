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
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gomodule/redigo/redis"

	"github.com/twitchtv/twirp"

	"github.com/sirupsen/logrus"
)

// hats
func main() {

	var hatRepo repo.HatRepo

	if config.RedisAddress != "" {
		conn, err := redis.Dial("tcp", config.RedisAddress)
		if err != nil {
			panic(fmt.Sprintf("failed to dial redis connection at %s err=%#v", config.RedisAddress, err))
		}
		defer conn.Close()
		redisrepo.NewRepo(conn)
	} else {
		hatRepo = inmem.NewRepo()
	}

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
