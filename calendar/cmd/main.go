package main

import (
	"calendar/internal/router"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	server := router.New(
		&http.ServeMux{},
		logger,
	)

	server.Run()
}
