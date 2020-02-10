package config

import (
	"backend/pkg/envconfig"
)

var ListenAddress = ":8083"
var MinSizeInches = int32(5)
var MaxSizeInches = int32(15)
var AppName = "hats"
var RedisAddress = ":6379"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetInt32("MIN_SIZE_INCHES", &MinSizeInches)
	envconfig.SetInt32("MAX_SIZE_INCHES", &MaxSizeInches)
	envconfig.SetString("REDIS_ADDRESS", &RedisAddress)
}
