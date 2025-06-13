package usecase

import (
	"context"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type CreateEventUseCase struct {
	repository storage.Storage //nolint:unused
	logger     logger.Logger   //nolint:unused
}

func NewCreateEventUseCase(repository storage.Storage, logger logger.Logger) *CreateEventUseCase {
	return &CreateEventUseCase{
		repository: repository,
		logger:     logger,
	}
}

func (c *CreateEventUseCase) Execute(_ context.Context, _ requestdto.CreateEvent) error {

	return nil
}
