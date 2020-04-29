package authnz

import (
	"backend/pkg/envconfig"
)

// Auth0AuthorizeURL is static here
var Auth0AuthorizeURL = "https://scorpicode.auth0.com/authorize"

// Auth0Issuer is static here
var Auth0Issuer = "https://scorpicode.auth0.com/"

// Auth0OAuthTokenURL is static here
var Auth0OAuthTokenURL = "https://scorpicode.auth0.com/oauth/token"

// Auth0Audience is static here
var Auth0Audience = "https://scorpicode.auth0.com/api/v2/"

// Auth0ResponseType is static here
var Auth0ResponseType = "code"

// Auth0ClientID always injected
var Auth0ClientID = ""

// Auth0ClientSecret always injected
var Auth0ClientSecret = ""

// Auth0RedirectURI local default but injected on Kubernetes
var Auth0RedirectURI = "http://localhost:8080/callback"

// Auth0PemfilePath local default but injected on Kubernetes
var Auth0PemfilePath = "../../scorpicode-local.pem"

// JWTExpirationHours hours generated tokens are valid; dev only, atm
var JWTExpirationHours = 24 * 365

func init() {
	envconfig.SetString("AUTH0_CLIENT_ID", &Auth0ClientID)
	envconfig.SetString("AUTH0_CLIENT_SECRET", &Auth0ClientSecret)
	envconfig.SetString("AUTH0_REDIRECT_URI", &Auth0RedirectURI)
	envconfig.SetString("AUTH0_PEMFILE_PATH", &Auth0PemfilePath)
}
