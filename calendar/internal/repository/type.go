package repository

import (
	"calendar/pkg/models"
	"log"
	"sync"
)

type Repository struct {
	events     map[uint]models.Event
	userEvents map[uint][]uint
	logger     *log.Logger
	mu         sync.RWMutex
}

func New(logger *log.Logger) *Repository {
	return &Repository{
		events:     make(map[uint]models.Event),
		userEvents: make(map[uint][]uint),
		logger:     logger,
	}
}
