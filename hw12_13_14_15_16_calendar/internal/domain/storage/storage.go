package storage

import (
	"context"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/entity"
)

type Storage interface {
	CreateEvent(ctx context.Context, event entity.Event) error
	UpdateEvent(ctx context.Context, event entity.Event) error
	DeleteEvent(ctx context.Context, eventID string) error
	ListEvents(ctx context.Context, userID string, from time.Time, to time.Time) ([]entity.Event, error)
	GetEvent(ctx context.Context, userID, eventID string) (entity.Event, error)
}
