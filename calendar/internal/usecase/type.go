// Package usecase содержит бизнес-логику приложения и координирует работу контроллеров и репозитория
package usecase

import (
	"calendar/internal/repository"
	"sync"

	"github.com/sirupsen/logrus"
)

// Service реализует бизнес-логику для работы с событиями календаря
type Service struct {
	repo   *repository.Repository
	logger *logrus.Logger
	nextID uint
	mu     sync.RWMutex
}

// New создаёт и инициализирует новый экземпляр Service
func New(repo *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}
