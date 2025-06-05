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

func NewCreateEvent(
	userID string,
	title string,
	description string,
	start time.Time,
	end time.Time,
	notifyBefore time.Duration,
) *CreateEvent {
	return &CreateEvent{
		UserID:       userID,
		Title:        title,
		Description:  description,
		TimeStart:    start,
		TimeEnd:      end,
		NotifyBefore: notifyBefore,
	}
}
