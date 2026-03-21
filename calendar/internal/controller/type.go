package controller

import (
	"calendar/internal/usecase"
	"log"
)

type Handler struct {
	service *usecase.Service
	logger  *log.Logger
}

func New(service *usecase.Service, logger *log.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
