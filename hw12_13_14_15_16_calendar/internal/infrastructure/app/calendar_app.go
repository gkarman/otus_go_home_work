package app

import (
	"context"
	"fmt"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/responsedto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/usecase"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type CalendarApp struct {
	logg logger.Logger
	st   storage.Storage
}

func NewCalendarApp(logger logger.Logger, st storage.Storage) *CalendarApp {
	return &CalendarApp{
		logg: logger,
		st:   st,
	}
}

func (app *CalendarApp) CreateEvent(ctx context.Context, req requestdto.CreateEvent) (responsedto.CreateEvent, error) {
	useCase := usecase.NewCreateEventUseCase(app.st, app.logg)
	response, err := useCase.Execute(ctx, req)
	if err != nil {
		return response, fmt.Errorf("calendar %w", err)
	}

	return response, nil
}
