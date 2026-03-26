package usecase

import (
	"calendar/pkg/models"
	"fmt"
	"strings"
	"time"
)

func EventErrorCheck(event *models.Event) error {
	if event.ID <= 0 {
		return fmt.Errorf("invalid event id")
	}
	if strings.TrimSpace(event.Title) == "" {
		return fmt.Errorf("invalid title")
	}
	if event.UserID <= 0 {
		return fmt.Errorf("invalid user_id")
	}
	return nil
}

func (s *Service) Create(eventReq *models.EventCreateRequest) (*models.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++

	date, err := time.Parse("2006-01-02", eventReq.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format (expected YYYY-MM-DD)")
	}

	event := &models.Event{
		ID:     s.nextID,
		UserID: eventReq.UserID,
		Title:  eventReq.Title,
		Date:   date,
	}
	if err := EventErrorCheck(event); err != nil {
		return nil, err
	}
	return s.repo.Create(event)
}

func (s *Service) Update(eventReq *models.EventUpdateRequest) (*models.Event, error) {
	date, err := time.Parse("2006-01-02", eventReq.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format (expected YYYY-MM-DD)")
	}

	event := &models.Event{
		ID:     eventReq.ID,
		UserID: eventReq.UserID,
		Title:  eventReq.Title,
		Date:   date,
	}
	if err := EventErrorCheck(event); err != nil {
		return nil, err
	}
	return s.repo.Update(event)
}
func (s *Service) Delete(eventId uint) error {
	if eventId <= 0 {
		return fmt.Errorf("invalid event id")
	}
	return s.repo.Delete(eventId)
}

func (s *Service) GetForDay(userID uint, date time.Time) ([]models.Event, error) {
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user_id")
	}
	return s.repo.GetForDay(userID, date)
}

func (s *Service) GetForWeek(userID uint, date time.Time) ([]models.Event, error) {
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user_id")
	}
	return s.repo.GetForWeek(userID, date)
}

func (s *Service) GetForMonth(userID uint, date time.Time) ([]models.Event, error) {
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user_id")
	}
	return s.repo.GetForMonth(userID, date)
}
