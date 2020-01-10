package config

import (
	"backend/pkg/envconfig"

	_ "github.com/sirupsen/logrus" // temp
)

var ListenAddress = ":8082"
var MinSizeInches = 0
var AppName = "hats"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetInt("MIN_SIZE_INCHES", &MinSizeInches)
}
