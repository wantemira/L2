package usecase

import (
	"calendar/internal/repository"
	"sync"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repo   *repository.Repository
	logger *logrus.Logger
	nextID uint
	mu     sync.RWMutex
}

func New(repo *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}
