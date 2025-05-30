package storage

import "time"

type Event struct {
	ID           string
	UserId       string
	Title        string
	TimeStart    time.Time
	TimeEnd      time.Time
	Description  string
	NotifyBefore time.Duration
}
