package config

import (
	"backend/pkg/envconfig"
)

// ListenAddress .
var ListenAddress = ":8084"

// AppName .
var AppName = "soxie"

// WriteWaitSeconds .
var WriteWaitSeconds = int64(10)

// PongWaitSeconds .
var PongWaitSeconds = int64(60)

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetInt64("WRITE_WAIT_SECONDS", &WriteWaitSeconds)
	envconfig.SetInt64("PONG_WAIT_SECONDS", &PongWaitSeconds)
}
