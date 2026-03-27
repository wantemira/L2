package repository

import (
	"calendar/pkg/models"
	"fmt"
	"time"
)

// Create добавляет новое событие в репозиторий
// Возвращает ошибку, если событие с таким ID уже существует
func (r *Repository) Create(event *models.Event) (*models.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.events[event.ID]; exists {
		return nil, fmt.Errorf("event already exists")
	}

	r.events[event.ID] = *event

	r.userEvents[event.UserID] = append(r.userEvents[event.UserID], event.ID)

	return event, nil
}

// Update обновляет существующее событие по его ID
// Возвращает ошибку, если событие не найдено
func (r *Repository) Update(event *models.Event) (*models.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.events[event.ID]; !exists {
		return nil, fmt.Errorf("event not found")
	}

	r.events[event.ID] = *event

	return event, nil
}

// Delete удаляет событие по его ID
// Возвращает ошибку, если событие не найдено
func (r *Repository) Delete(eventId uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	event, exists := r.events[eventId]
	if !exists {
		return fmt.Errorf("event not found")
	}

	delete(r.events, eventId)

	userEvents := r.userEvents[event.UserID]
	for i, id := range userEvents {
		if id == eventId {
			r.userEvents[event.UserID] = append(userEvents[:i], userEvents[i+1:]...)
			break
		}
	}

	return nil
}

// GetForDay возвращает все события пользователя за указанный день
func (r *Repository) GetForDay(userID uint, date time.Time) ([]models.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]models.Event, 0)
	eventIDs, exists := r.userEvents[userID]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	for _, id := range eventIDs {
		event := r.events[id]
		if isDay(event.Date, date) {
			result = append(result, event)
		}
	}

	return result, nil
}

// GetForWeek возвращает все события пользователя за неделю, содержащую указанную дату
func (r *Repository) GetForWeek(userID uint, date time.Time) ([]models.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]models.Event, 0)

	eventIDs, exists := r.userEvents[userID]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	startWeek, endWeek := getRangeOfWeek(date)

	for _, id := range eventIDs {
		event := r.events[id]
		if (event.Date.Equal(startWeek) || event.Date.After(startWeek)) && event.Date.Before(endWeek) {
			result = append(result, event)
		}
	}

	return result, nil
}

// GetForMonth возвращает все события пользователя за месяц, содержащий указанную дату
func (r *Repository) GetForMonth(userID uint, date time.Time) ([]models.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]models.Event, 0)

	eventIDs, exists := r.userEvents[userID]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	for _, id := range eventIDs {
		event := r.events[id]
		if isInMonth(event.Date, date) && event.Date.AddDate(0, 0, 1).Compare(date) > 0 {
			result = append(result, event)
		}
	}

	return result, nil
}

func isInMonth(day, expected time.Time) bool {
	y1, m1, _ := day.Date()
	y2, m2, _ := expected.Date()
	return y1 == y2 && m1 == m2
}

func getRangeOfWeek(day time.Time) (time.Time, time.Time) {
	weekday := day.Weekday()
	offset := int(weekday)
	if offset == 0 {
		offset = 7
	}
	startWeek := day.AddDate(0, 0, -offset+1)
	return startWeek, startWeek.AddDate(0, 0, 7)
}

func isDay(day, expected time.Time) bool {
	y1, m1, d1 := day.Date()
	y2, m2, d2 := expected.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
