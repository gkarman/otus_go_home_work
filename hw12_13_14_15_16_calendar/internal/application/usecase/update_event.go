package usecase

import (
	"context"
	"fmt"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type UpdateEventUseCase struct {
	st     storage.Storage
	logger logger.Logger
}

func NewUpdateEventUseCase(st storage.Storage, logger logger.Logger) *UpdateEventUseCase {
	return &UpdateEventUseCase{
		st:     st,
		logger: logger,
	}
}

func (c *UpdateEventUseCase) Execute(ctx context.Context, req *requestdto.UpdateEvent) error {
	event, err := c.st.GetEvent(ctx, req.ID, req.UserID)
	if err != nil {
		return fmt.Errorf("get for update in storage: %w", err)
	}

	event.Title = req.Title
	event.Description = req.Description
	event.TimeStart = req.TimeStart
	event.TimeEnd = req.TimeEnd
	event.NotifyBefore = req.NotifyBefore

	err = c.st.UpdateEvent(ctx, event)
	if err != nil {
		return fmt.Errorf("update in storage %w", err)
	}

	return nil
}
