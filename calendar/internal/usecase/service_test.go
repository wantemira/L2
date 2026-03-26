package usecase

import (
	"calendar/internal/repository"
	"calendar/pkg/models"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func setupService(t *testing.T) *Service {
	t.Helper()
	logger := logrus.New()
	repo := repository.New(logger)
	return New(repo, logger)
}

func createTestEvent(t *testing.T, s *Service, date string) *models.Event {
	t.Helper()
	req := &models.EventCreateRequest{
		UserID: 1,
		Title:  "Test Event",
		Date:   date,
	}
	event, err := s.Create(req)
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	return event
}

func TestService_Create(t *testing.T) {
	s := setupService(t)

	req := &models.EventCreateRequest{
		UserID: 1,
		Title:  "New Event",
		Date:   "2023-12-25",
	}

	event, err := s.Create(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event.ID != 1 {
		t.Errorf("expected ID 1, got %d", event.ID)
	}
}

func TestService_GetForDay(t *testing.T) {
	s := setupService(t)

	created := createTestEvent(t, s, "2023-12-25")

	date := parseDate("2023-12-25")
	events, err := s.GetForDay(1, date)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("expected 1 event, got %d", len(events))
	}
	if events[0].ID != created.ID {
		t.Errorf("expected event ID %d, got %d", created.ID, events[0].ID)
	}
}

func TestService_Delete(t *testing.T) {
	s := setupService(t)

	created := createTestEvent(t, s, "2023-12-25")

	err := s.Delete(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	events, _ := s.GetForDay(1, parseDate("2023-12-25"))
	if len(events) != 0 {
		t.Errorf("expected 0 events after delete, got %d", len(events))
	}
}

func parseDate(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}
