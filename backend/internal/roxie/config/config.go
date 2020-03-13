package config

import (
	"backend/pkg/envconfig"
)

// ListenAddress .
var ListenAddress = ":8080"

// AppName .
var AppName = "roxie"

// WebsitePrefix .
var WebsitePrefix = "http://localhost:8081"

// FrontendPrefix .
var FrontendPrefix = "http://localhost:3000"

// HatsPrefix .
var HatsPrefix = "http://localhost:8083"

// EnableCors .
var EnableCors = true

// LoginSuccessTarget .
var LoginSuccessTarget = "http://localhost:8080/sc/"

// Auth0AuthorizeURL is static here
var Auth0AuthorizeURL = "https://scorpicode.auth0.com/authorize"

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

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetString("WEBSITE_PREFIX", &WebsitePrefix)
	envconfig.SetString("FRONTEND_PREFIX", &FrontendPrefix)
	envconfig.SetString("HATS_PREFIX", &HatsPrefix)
	envconfig.SetBool("ENABLE_CORS", &EnableCors)

	envconfig.SetString("LOGIN_SUCCESS_TARGET", &LoginSuccessTarget)
	envconfig.SetString("AUTH0_CLIENT_ID", &Auth0ClientID)
	envconfig.SetString("AUTH0_CLIENT_SECRET", &Auth0ClientSecret)
	envconfig.SetString("AUTH0_REDIRECT_URI", &Auth0RedirectURI)
}
