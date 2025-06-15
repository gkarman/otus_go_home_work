package usecase

import (
	"context"
	"fmt"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/responsedto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/storage"
	"github.com/google/uuid"
)

type CreateEventUseCase struct {
	st     storage.Storage
	logger logger.Logger
}

func NewCreateEventUseCase(st storage.Storage, logger logger.Logger) *CreateEventUseCase {
	return &CreateEventUseCase{
		st:     st,
		logger: logger,
	}
}

func (c *CreateEventUseCase) Execute(_ context.Context, req requestdto.CreateEvent) (responsedto.CreateEvent, error) {
	eventID := uuid.New()
	event := entity.Event{
		ID:           eventID.String(),
		UserID:       req.UserID,
		Title:        req.Title,
		Description:  req.Description,
		TimeStart:    req.TimeStart,
		TimeEnd:      req.TimeEnd,
		NotifyBefore: req.NotifyBefore,
	}

	err := c.st.CreateEvent(context.Background(), event)
	if err != nil {
		return responsedto.CreateEvent{}, fmt.Errorf("save in storage %w", err)
	}

	response := responsedto.CreateEvent{
		ID: eventID.String(),
	}

	return response, nil
}
