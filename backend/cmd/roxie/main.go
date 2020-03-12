package main

import (
	"backend/internal/roxie/config"
	"backend/internal/roxie/server"
	_ "backend/pkg/logging" // init logrus
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// roxie
func main() {

	logrus.Infof("main() %s", config.AppName)

	proxies := []*server.Proxy{
		// website
		&server.Proxy{
			HostPrefix: config.WebsitePrefix,
			FromPath:   "/",
			ToPath:     "/",
		},
		&server.Proxy{
			HostPrefix: config.WebsitePrefix,
			FromPath:   "/login",
			ToPath:     "/login",
		},
		&server.Proxy{
			HostPrefix: config.WebsitePrefix,
			FromPath:   "/logincallback",
			ToPath:     "/logincallback",
		},
		&server.Proxy{
			HostPrefix: config.WebsitePrefix,
			FromPath:   "/logout",
			ToPath:     "/logout",
		},

		// frontend
		&server.Proxy{
			HostPrefix: config.FrontendPrefix,
			FromPath:   "/sc/",
			ToPath:     "/",
		},
		// hats
		&server.Proxy{
			HostPrefix: config.HatsPrefix,
			FromPath:   "/hats/",
			ToPath:     "/twirp/hats.Hats/",
		},
	}

	mux := http.NewServeMux()

	for _, p := range proxies {
		mux.HandleFunc(p.FromPath, p.GetHandle())

		logrus.Infof("main() %s proxying %s%s => %s%s", config.AppName, config.ListenAddress, p.FromPath, p.HostPrefix, p.ToPath)
	}

	logrus.Infof("main() %s listening on %s", config.AppName, config.ListenAddress)

	GracefulServe(config.ListenAddress, mux)
}

// GracefulServe .
func GracefulServe(addr string, mux *http.ServeMux) {

	if len(addr) == 0 {
		logrus.Fatal("GracefulServe() no address specified")
	}

	s := http.Server{Addr: addr, Handler: mux}

	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
		s.Shutdown(ctxShutDown)
	})

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(err)
	}
	logrus.Infof("exited gracefully")
}
