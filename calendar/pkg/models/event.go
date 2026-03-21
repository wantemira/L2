package models

import "time"

type Event struct {
	ID     uint
	UserID uint
	Title  string
	Date   time.Time
}

type EventRequest struct {
	UserID uint      `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}
