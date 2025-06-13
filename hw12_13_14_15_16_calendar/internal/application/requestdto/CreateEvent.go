package requestdto

import "time"

type CreateEvent struct {
	UserID       string
	Title        string
	Description  string
	TimeStart    time.Time
	TimeEnd      time.Time
	NotifyBefore time.Duration
}
