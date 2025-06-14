package usecase

import (
	"context"
	"fmt"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type DeleteEventUseCase struct {
	st     storage.Storage
	logger logger.Logger
}

func NewDeleteEventUseCase(st storage.Storage, logger logger.Logger) *DeleteEventUseCase {
	return &DeleteEventUseCase{
		st:     st,
		logger: logger,
	}
}

func (c *DeleteEventUseCase) Execute(ctx context.Context, req *requestdto.DeleteEvent) error {
	event, err := c.st.GetEvent(ctx, req.ID, req.UserID)
	if err != nil {
		return fmt.Errorf("get in storage %w", err)
	}

	err = c.st.DeleteEvent(ctx, event.ID)
	if err != nil {
		return fmt.Errorf("delete in storage: %w", err)
	}

	return nil
}
