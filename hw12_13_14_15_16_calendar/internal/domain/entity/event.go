package entity

import "time"

type Event struct {
	ID           string        `db:"id"`
	UserID       string        `db:"user_id"`
	Title        string        `db:"title"`
	Description  string        `db:"description"`
	TimeStart    time.Time     `db:"time_start"`
	TimeEnd      time.Time     `db:"time_end"`
	NotifyBefore time.Duration `db:"notify_before"`
}
