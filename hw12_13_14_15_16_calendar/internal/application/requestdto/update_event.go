package requestdto

import "time"

type UpdateEvent struct {
	ID           string        `json:"id"`
	UserID       string        `json:"userId"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	TimeStart    time.Time     `json:"timeStart"`
	TimeEnd      time.Time     `json:"timeEnd"`
	NotifyBefore time.Duration `json:"notifyBefore"`
}
