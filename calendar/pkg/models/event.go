// Package models содержит определения структур данных для работы с событиями календаря
package models

import "time"

// Event представляет событие в календаре пользователя
type Event struct {
	ID     uint
	UserID uint
	Title  string
	Date   time.Time
}

// EventCreateRequest содержит данные для создания нового события
type EventCreateRequest struct {
	UserID uint   `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}

// EventUpdateRequest содержит данные для обновления существующего события
type EventUpdateRequest struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}
