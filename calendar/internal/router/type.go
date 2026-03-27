// Package router настраивает маршрутизацию HTTP-запросов и запускает сервер
package router

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// Server представляет HTTP-сервер с настроенными маршрутами и зависимостями
type Server struct {
	mux    *http.ServeMux
	logger *logrus.Logger
}

// New создаёт и инициализирует новый экземпляр Server
func New(mux *http.ServeMux,
	logger *logrus.Logger) *Server {
	return &Server{
		mux:    mux,
		logger: logger,
	}
}
