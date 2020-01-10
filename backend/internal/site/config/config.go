package config

import (
	"backend/pkg/envconfig"
)

var ListenAddress = ":8081"
var AppName = "site"
var StaticPath = "./static"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetString("STATIC_PATH", &StaticPath)
}
