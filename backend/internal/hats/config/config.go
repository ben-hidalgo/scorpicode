package config

import (
	"backend/pkg/envconfig"
)

// ListenAddress .
var ListenAddress = ":8083"

// MinSizeInches .
var MinSizeInches = int32(5)

// MaxSizeInches .
var MaxSizeInches = int32(15)

// AppName .
var AppName = "hats"

// DatastoreConfig .
var DatastoreConfig = "inmem"

func init() {
	envconfig.SetString("LISTEN_ADDRESS", &ListenAddress)
	envconfig.SetString("APP_NAME", &AppName)
	envconfig.SetInt32("MIN_SIZE_INCHES", &MinSizeInches)
	envconfig.SetInt32("MAX_SIZE_INCHES", &MaxSizeInches)
	envconfig.SetString("DATASTORE_CONFIG", &DatastoreConfig)
}
