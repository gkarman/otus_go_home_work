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

func (app *CalendarApp) DeleteEvent(ctx context.Context, req requestdto.DeleteEvent) error {
	useCase := usecase.NewDeleteEventUseCase(app.st, app.logg)
	err := useCase.Execute(ctx, &req)
	if err != nil {
		return fmt.Errorf("calendar %w", err)
	}

	return nil
}

func (app *CalendarApp) UpdateEvent(ctx context.Context, req requestdto.UpdateEvent) error {
	useCase := usecase.NewUpdateEventUseCase(app.st, app.logg)
	err := useCase.Execute(ctx, &req)
	if err != nil {
		return fmt.Errorf("calendar %w", err)
	}

	return nil
}

func (app *CalendarApp) EventsDay(ctx context.Context, req requestdto.EventsOnDate) (*responsedto.Events, error) {
	useCase := usecase.NewEventsDayUseCase(app.st, app.logg)
	events, err := useCase.Execute(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("calendar %w", err)
	}

	resp := responsedto.Events{
		Events: events,
	}

	return &resp, nil
}

func (app *CalendarApp) EventsWeek(ctx context.Context, req requestdto.EventsOnDate) (*responsedto.Events, error) {
	useCase := usecase.NewEventsWeekUseCase(app.st, app.logg)
	events, err := useCase.Execute(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("calendar %w", err)
	}

	resp := responsedto.Events{
		Events: events,
	}

	return &resp, nil
}

func (app *CalendarApp) EventsMonth(ctx context.Context, req requestdto.EventsOnDate) (*responsedto.Events, error) {
	useCase := usecase.NewEventsMonthUseCase(app.st, app.logg)
	events, err := useCase.Execute(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("calendar %w", err)
	}

	resp := responsedto.Events{
		Events: events,
	}

	return &resp, nil
}
