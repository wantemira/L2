package usecase

import (
	"calendar/internal/repository"
	"log"
	"sync"
)

type Service struct {
	repo   *repository.Repository
	logger *log.Logger
	nextID uint
	mu     sync.RWMutex
}

func New(repo *repository.Repository, logger *log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}
