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

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetString("STATIC_PATH", &StaticPath)
}
