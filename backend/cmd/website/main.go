package main

import (
	"backend/internal/site/config"
	_ "backend/pkg/logging" // init logrus
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(config.StaticPath)))

	logrus.Infof("main() %s listening on %s", config.AppName, config.ListenAddress)

	logrus.Fatal(http.ListenAndServe(config.ListenAddress, nil))
}
