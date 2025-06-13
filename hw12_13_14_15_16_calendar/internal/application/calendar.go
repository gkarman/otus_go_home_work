package application

import (
	"context"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
)

type Calendar interface {
	CreateEvent(ctx context.Context, event requestdto.CreateEvent) error
}
