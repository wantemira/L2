package models

import "time"

type Event struct {
	ID     uint
	UserID uint
	Title  string
	Date   time.Time
}
