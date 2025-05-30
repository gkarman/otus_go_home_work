package storage

import (
	"context"
	"time"
)

type EventStorage interface {
	AddEvent(ctx context.Context, event Event) error
	UpdateEvent(ctx context.Context, event Event) error
	DeleteEvent(ctx context.Context, eventID string) error
	ListEvents(ctx context.Context, userID string, from time.Time, to time.Time) ([]Event, error)
}
