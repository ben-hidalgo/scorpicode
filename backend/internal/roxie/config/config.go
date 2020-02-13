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

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetString("WEBSITE_PREFIX", &WebsitePrefix)
	envconfig.SetString("FRONTEND_PREFIX", &FrontendPrefix)
	envconfig.SetString("HATS_PREFIX", &HatsPrefix)

	envconfig.SetBool("ENABLE_CORS", &EnableCors)
}
