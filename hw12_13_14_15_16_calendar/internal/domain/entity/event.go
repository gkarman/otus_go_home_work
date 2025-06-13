package entity

import "time"

type Event struct {
	ID           string
	UserID       string
	Title        string
	Description  string
	TimeStart    time.Time
	TimeEnd      time.Time
	NotifyBefore time.Duration
}
