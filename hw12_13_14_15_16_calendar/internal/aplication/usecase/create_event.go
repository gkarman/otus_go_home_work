package usecase

import (
	"context"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/aplication/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type CreateEvent struct {
	repository storage.Storage
	logger     logger.Logger
}

func (c *CreateEvent) Run(ctx context.Context, request requestdto.CreateEvent) error {
	return nil
}
