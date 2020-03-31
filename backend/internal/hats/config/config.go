package config

import (
	"backend/pkg/envconfig"
)

// ListenAddress .
var ListenAddress = ":8083"

// AppName .
var AppName = "hats"

// DatastoreConfig .
var DatastoreConfig = "inmem"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetString("DATASTORE_CONFIG", &DatastoreConfig)
}
