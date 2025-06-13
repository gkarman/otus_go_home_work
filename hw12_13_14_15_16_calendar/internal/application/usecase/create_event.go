package usecase

import (
	"context"
	"fmt"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type CreateEventUseCase struct {
	st     storage.Storage //nolint:unused
	logger logger.Logger   //nolint:unused
}

func NewCreateEventUseCase(st storage.Storage, logger logger.Logger) *CreateEventUseCase {
	return &CreateEventUseCase{
		st:     st,
		logger: logger,
	}
}

func (c *CreateEventUseCase) Execute(_ context.Context, request requestdto.CreateEvent) error {

	event := entity.Event{
		"11111111-1111-1111-1111-111111111112",
		request.UserID,
		request.Title,
		request.Description,
		request.TimeStart,
		request.TimeEnd,
		request.NotifyBefore,
	}

	err := c.st.CreateEvent(context.Background(), event)
	if err != nil {
		return fmt.Errorf("save in storage %w", err)
	}

	return nil
}
