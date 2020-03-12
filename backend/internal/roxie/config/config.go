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
var FrontendPrefix = "http://localhost:8082"

// HatsPrefix .
var HatsPrefix = "http://localhost:8083"

// EnableCors .
var EnableCors = true

// LoginSuccessTarget .
var LoginSuccessTarget = "http://localhost:3000"

// Auth0AuthorizeURL .
var Auth0AuthorizeURL = "https://scorpicode.auth0.com/authorize"

// Auth0ResponseType .
//var Auth0ResponseType = "token"
var Auth0ResponseType = "code"

// Auth0ClientID .
var Auth0ClientID = ""

// Auth0RedirectURI .
var Auth0RedirectURI = "http://localhost:8080/callback/"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetString("WEBSITE_PREFIX", &WebsitePrefix)
	envconfig.SetString("FRONTEND_PREFIX", &FrontendPrefix)
	envconfig.SetString("HATS_PREFIX", &HatsPrefix)
	envconfig.SetBool("ENABLE_CORS", &EnableCors)

	envconfig.SetString("LOGIN_SUCCESS_TARGET", &LoginSuccessTarget)
	envconfig.SetString("AUTH0_AUTHORIZE_URL", &Auth0AuthorizeURL)
	envconfig.SetString("AUTH0_RESPONSE_TYPE", &Auth0ResponseType)
	envconfig.SetString("AUTH0_CLIENT_ID", &Auth0ClientID)
	envconfig.SetString("AUTH0_REDIRECT_URI", &Auth0RedirectURI)
}
