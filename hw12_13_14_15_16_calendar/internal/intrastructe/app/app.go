package app

import (
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type App struct {
	logger     logger.Logger
	repository storage.Storage
}

func New(logger logger.Logger, repository storage.Storage) *App {
	return &App{
		logger:     logger,
		repository: repository,
	}
}
