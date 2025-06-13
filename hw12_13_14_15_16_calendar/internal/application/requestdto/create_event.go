package requestdto

import "time"

type CreateEvent struct {
	UserID       string        `json:"userId"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	TimeStart    time.Time     `json:"timeStart"`
	TimeEnd      time.Time     `json:"timeEnd"`
	NotifyBefore time.Duration `json:"notifyBefore"`
}
