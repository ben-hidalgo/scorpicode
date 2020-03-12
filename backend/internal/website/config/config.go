package config

import (
	"backend/pkg/envconfig"
)

// ListenAddress .
var ListenAddress = ":8081"

// AppName .
var AppName = "site"

// StaticPath .
var StaticPath = "./static"

// Auth0AuthorizeURL .
var Auth0AuthorizeURL = "https://scorpicode.auth0.com/authorize"

// Auth0ResponseType .
var Auth0ResponseType = "token"

// Auth0ClientID .
var Auth0ClientID = ""

// Auth0RedirectURI .
var Auth0RedirectURI = "http://localhost:8080/logincallback"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetString("STATIC_PATH", &StaticPath)

	envconfig.SetString("AUTH0_AUTHORIZE_URL", &Auth0AuthorizeURL)
	envconfig.SetString("AUTH0_RESPONSE_TYPE", &Auth0ResponseType)
	envconfig.SetString("AUTH0_CLIENT_ID", &Auth0ClientID)
	envconfig.SetString("AUTH0_REDIRECT_URI", &Auth0RedirectURI)
}
