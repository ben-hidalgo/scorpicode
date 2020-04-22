package endpoints

import (
	"backend/internal/roxie/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Callback handles the Auth0 callback
func Callback(w http.ResponseWriter, r *http.Request) {

	logrus.Trace("Callback() received")

	code := r.URL.Query().Get("code")

	logrus.Tracef("callback() r.URL.Query()=%#v", r.URL.Query())

	reqBody, err := json.Marshal(map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     config.Auth0ClientID,
		"client_secret": config.Auth0ClientSecret,
		"code":          code,
		"redirect_uri":  config.Auth0RedirectURI,
		"audience":      config.Auth0Audience,
	})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := http.Post(config.Auth0OAuthTokenURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	logrus.Tracef("callback() oauth/token response=%s", b)

	var dat map[string]interface{}

	if err := json.Unmarshal(b, &dat); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO: conditionally on env var, use secure cookie document.cookie = "name = value;secure";
	// TODO: make the cookie name environment specific or switch to local storage
	tmpl := `
	<!DOCTYPE html> 
	<html>
	<body> 
		<script>
			function setCookie(name, value, days) {
				var d = new Date;
				d.setTime(d.getTime() + 24*60*60*1000*days);
				// not using string interpolation because we're inside a Go string literal
				document.cookie = name + "=" + value + ";path=/;expires=" + d.toGMTString();
			}
			window.onload = () => { 
				setCookie('id_token', '%s', 3)
				window.location.href = '%s'
			} 
		</script> 
	</body> 
	</html>
	`
	token := dat["id_token"]

	if config.LocalHeadersPath != "" {
		writeLocalHeaders(config.LocalHeadersPath, fmt.Sprint(token))
	}

	body := fmt.Sprintf(tmpl, token, config.LoginSuccessTarget)

	w.WriteHeader(200)
	w.Write([]byte(body))
	return
}

func writeLocalHeaders(path, token string) {

	tmpl := `content-type: application/json
authorization: Bearer %s
`

	body := fmt.Sprintf(tmpl, token)

	ioutil.WriteFile(path, []byte(body), 0644)

}
