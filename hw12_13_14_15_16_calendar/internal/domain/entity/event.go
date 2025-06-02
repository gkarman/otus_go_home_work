package entity

import "time"

type Event struct {
	ID           string
	UserID       string
	Title        string
	TimeStart    time.Time
	TimeEnd      time.Time
	Description  string
	NotifyBefore time.Duration
}
