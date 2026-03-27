// Package repository предоставляет слой доступа к данным для хранения событий в памяти
package repository

import (
	"calendar/pkg/models"
	"sync"

	"github.com/sirupsen/logrus"
)

// Repository хранит события в памяти и обеспечивает потокобезопасный доступ к ним
type Repository struct {
	events     map[uint]models.Event
	userEvents map[uint][]uint
	logger     *logrus.Logger
	mu         sync.RWMutex
}

// New создаёт и инициализирует новый экземпляр Repository
func New(logger *logrus.Logger) *Repository {
	return &Repository{
		events:     make(map[uint]models.Event),
		userEvents: make(map[uint][]uint),
		logger:     logger,
	}
}
