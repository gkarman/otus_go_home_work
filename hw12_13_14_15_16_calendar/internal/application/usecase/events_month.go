package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
)

type EventsMonthUseCase struct {
	st     storage.Storage
	logger logger.Logger
}

func NewEventsMonthUseCase(st storage.Storage, logger logger.Logger) *EventsMonthUseCase {
	return &EventsMonthUseCase{
		st:     st,
		logger: logger,
	}
}

func (c *EventsMonthUseCase) Execute(ctx context.Context, req *requestdto.EventsOnDate) ([]entity.Event, error) {
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}
	from := time.Date(parsedDate.Year(), parsedDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	to := from.AddDate(0, 1, 0)

	result, err := c.st.ListEvents(ctx, req.UserID, from, to)
	if err != nil {
		return nil, fmt.Errorf("get list in storage %w", err)
	}
	return result, nil
}
