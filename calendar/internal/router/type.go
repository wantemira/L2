package router

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	mux    *http.ServeMux
	logger *logrus.Logger
}

func New(mux *http.ServeMux,
	logger *logrus.Logger) *Server {
	return &Server{
		mux:    mux,
		logger: logger,
	}
}
