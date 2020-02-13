package logging

import (
	"backend/pkg/envconfig"
	"fmt"

	"github.com/sirupsen/logrus"
)

// LogLevel default value
var LogLevel = "debug"

// LogFormat default value
var LogFormat = "text"

func init() {
	envconfig.SetString("LOG_LEVEL", &LogLevel)

	level, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		panic(fmt.Sprintf("logging.init() could not parse log level=%s", LogLevel))
	}
	logrus.SetLevel(level)
}

func init() {
	envconfig.SetString("LOG_FORMAT", &LogFormat)

	switch LogFormat {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		panic(fmt.Sprintf("logging.init() unexpected format=%s", LogFormat))
	}
}
