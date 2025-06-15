package application

import (
	"context"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/responsedto"
)

type Calendar interface {
	CreateEvent(ctx context.Context, req requestdto.CreateEvent) (responsedto.CreateEvent, error)
	DeleteEvent(ctx context.Context, req requestdto.DeleteEvent) error
	UpdateEvent(ctx context.Context, req requestdto.UpdateEvent) error
	EventsDay(ctx context.Context, req requestdto.EventsOnDate) (*responsedto.Events, error)
	EventsWeek(ctx context.Context, req requestdto.EventsOnDate) (*responsedto.Events, error)
	EventsMonth(ctx context.Context, req requestdto.EventsOnDate) (*responsedto.Events, error)
}
