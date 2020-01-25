package config

import (
	"backend/pkg/envconfig"
)

var ListenAddress = ":8080"
var AppName = "roxie"

var WebsitePrefix = "http://localhost:8081"
var FrontendPrefix = "http://localhost:8082"
var HatsPrefix = "http://localhost:8083"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetString("WEBSITE_PREFIX", &WebsitePrefix)
	envconfig.SetString("FRONTEND_PREFIX", &FrontendPrefix)
	envconfig.SetString("HATS_PREFIX", &HatsPrefix)
}
