// Package controller обрабатывает HTTP-запросы и делегирует бизнес-логику слою usecase
package controller

import (
	"calendar/internal/usecase"

	"github.com/sirupsen/logrus"
)

// Handler обрабатывает HTTP-запросы к эндпоинтам календаря
type Handler struct {
	service *usecase.Service
	logger  *logrus.Logger
}

// New создаёт и инициализирует новый экземпляр Handler
func New(service *usecase.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
