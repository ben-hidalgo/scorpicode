package main

import (
	"backend/internal/website/config"
	_ "backend/pkg/logging" // init logrus
	"net/http"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

// website
func main() {

	http.Handle("/", http.FileServer(http.Dir(config.StaticPath)))

	http.HandleFunc("/login", login)

	logrus.Infof("main() %s listening on %s", config.AppName, config.ListenAddress)

	logrus.Fatal(http.ListenAndServe(config.ListenAddress, nil))
}

// TODO: this won't work in a clustered env...
// just sign a timestamp and check if it's signed and not old
// string uuid to boolean "is consumed"; will always be true
var states = map[string]bool{}

func login(w http.ResponseWriter, r *http.Request) {

	// using req to build the redirect URL
	req, err := http.NewRequest("GET", config.Auth0AuthorizeURL, nil)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	uuid := xid.New().String()

	states[uuid] = true

	q := req.URL.Query()
	q.Add("response_type", config.Auth0ResponseType)
	q.Add("client_id", config.Auth0ClientID)
	q.Add("redirect_uri", config.Auth0RedirectURI)
	q.Add("state", uuid)
	req.URL.RawQuery = q.Encode()

	http.Redirect(w, r, req.URL.String(), 302)
}
