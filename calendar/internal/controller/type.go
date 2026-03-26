package controller

import (
	"calendar/internal/usecase"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *usecase.Service
	logger  *logrus.Logger
}

func New(service *usecase.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
