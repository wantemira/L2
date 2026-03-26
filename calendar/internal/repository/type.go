package repository

import (
	"calendar/pkg/models"
	"sync"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	events     map[uint]models.Event
	userEvents map[uint][]uint
	logger     *logrus.Logger
	mu         sync.RWMutex
}

func New(logger *logrus.Logger) *Repository {
	return &Repository{
		events:     make(map[uint]models.Event),
		userEvents: make(map[uint][]uint),
		logger:     logger,
	}
}
