package app

import (
	"context"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type CalendarApp struct {
	logger     logger.Logger
	repository storage.Storage
}

func NewCalendarApp(logger logger.Logger, repository storage.Storage) *CalendarApp {
	return &CalendarApp{
		logger:     logger,
		repository: repository,
	}
}

func (app *CalendarApp) CreateEvent(_ context.Context, event requestdto.CreateEvent) error {
	return nil
}
