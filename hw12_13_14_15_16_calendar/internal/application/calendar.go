package application

import (
	"context"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/responsedto"
)

type Calendar interface {
	CreateEvent(ctx context.Context, event requestdto.CreateEvent) (responsedto.CreateEvent, error)
	DeleteEvent(ctx context.Context, event requestdto.DeleteEvent) error
}
