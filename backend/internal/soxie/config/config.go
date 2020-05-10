package config

import (
	"backend/pkg/envconfig"
)

// ListenAddress .
var ListenAddress = ":8084"

// AppName .
var AppName = "soxie"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
}
