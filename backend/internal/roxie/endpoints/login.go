package endpoints

import (
	"backend/internal/roxie/config"
	"net/http"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

// Login redirects to the Auth0 universal login
func Login(w http.ResponseWriter, r *http.Request) {

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
	q.Add("scope", "openid profile email")
	req.URL.RawQuery = q.Encode()

	logrus.Debugf("login() req.URL.Query()=%#v", req.URL.Query())

	http.Redirect(w, r, req.URL.String(), 302)
}
