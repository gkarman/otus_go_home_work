package storage

import (
	"context"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain"
)

type StorageInterface interface {
	CreateEvent(ctx context.Context, event domain.Event) error
	UpdateEvent(ctx context.Context, event domain.Event) error
	DeleteEvent(ctx context.Context, eventID string) error
	ListEvents(ctx context.Context, userID string, from time.Time, to time.Time) ([]domain.Event, error)
}
