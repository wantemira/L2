package models

import "time"

type Event struct {
	ID     uint
	UserID uint
	Title  string
	Date   time.Time
}

type EventCreateRequest struct {
	UserID uint   `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}
type EventUpdateRequest struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}
