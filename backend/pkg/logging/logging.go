package logging

import (
	"backend/pkg/envconfig"
	"fmt"

	"github.com/sirupsen/logrus"
)

var LogLevel = "DEBUG"

func init() {
	envconfig.SetString("LOG_LEVEL", &LogLevel)

	level, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		panic(fmt.Sprintf("logging init() could not parse log level=%s", LogLevel))
	}
	logrus.SetLevel(level)
}
