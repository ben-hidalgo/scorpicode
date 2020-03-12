package main

import (
	"backend/internal/roxie/config"
	"backend/internal/roxie/server"
	_ "backend/pkg/logging" // init logrus
	"net/http"

	"github.com/rs/xid"
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

	mux.HandleFunc("/login", login)

	mux.HandleFunc("/callback/", callback)

	logrus.Infof("main() %s listening on %s", config.AppName, config.ListenAddress)

	GracefulServe(config.ListenAddress, mux)
}

func callback(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("callback() r=%#v", r)
	logrus.Infof("callback() r.URL=%s", r.URL)
	logrus.Infof("callback() r.URL.Query()=%#v", r.URL.Query())

	w.WriteHeader(200)
	w.Write([]byte("Welcome!"))
	return

	// http.Redirect(w, r, "/sc", 302)
}

func login(w http.ResponseWriter, r *http.Request) {

	// using req to build the redirect URL
	req, err := http.NewRequest("GET", config.Auth0AuthorizeURL, nil)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO: sign a timestamp and check if it's signed and not old
	uuid := xid.New().String()

	q := req.URL.Query()
	q.Add("response_type", config.Auth0ResponseType)
	q.Add("client_id", config.Auth0ClientID)
	q.Add("redirect_uri", config.Auth0RedirectURI)
	q.Add("state", uuid)
	req.URL.RawQuery = q.Encode()

	http.Redirect(w, r, req.URL.String(), 302)
}

// GracefulServe .
func GracefulServe(addr string, mux *http.ServeMux) {

	if len(addr) == 0 {
		logrus.Fatal("GracefulServe() no address specified")
	}

	s := http.Server{Addr: addr, Handler: mux}

	// mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
	// 	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer func() {
	// 		cancel()
	// 	}()
	// 	s.Shutdown(ctxShutDown)
	// })

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(err)
	}
	logrus.Infof("exited gracefully")
}
