package memorystorage

import (
	"context"
	"testing"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func makeTestEvent(userID string) entity.Event {
	id := uuid.New().String()
	start := time.Now().Add(time.Hour)
	return entity.Event{
		ID:           id,
		Title:        "Test Event",
		TimeStart:    start,
		TimeEnd:      start.Add(time.Hour),
		Description:  "A test event",
		UserId:       userID,
		NotifyBefore: 10 * time.Minute,
	}
}

func TestStorage_CreateEvent(t *testing.T) {
	s := New()
	ctx := context.Background()

	event := makeTestEvent("user1")

	err := s.CreateEvent(ctx, event)
	require.NoError(t, err)

	err = s.CreateEvent(ctx, event)
	require.ErrorIs(t, err, domain.ErrEntityAlreadyExists)
}

func TestStorage_UpdateEvent(t *testing.T) {
	s := New()
	ctx := context.Background()

	event := makeTestEvent("user1")
	err := s.CreateEvent(ctx, event)
	require.NoError(t, err)

	event.Title = "Updated title"
	err = s.UpdateEvent(ctx, event)
	require.NoError(t, err)

	nonexistent := makeTestEvent("user1")
	err = s.UpdateEvent(ctx, nonexistent)
	require.ErrorIs(t, err, domain.ErrEntityNotFound)
}

func TestStorage_DeleteEvent(t *testing.T) {
	s := New()
	ctx := context.Background()

	event := makeTestEvent("user1")
	err := s.CreateEvent(ctx, event)
	require.NoError(t, err)

	err = s.DeleteEvent(ctx, event.ID)
	require.NoError(t, err)

	err = s.DeleteEvent(ctx, event.ID)
	require.ErrorIs(t, err, domain.ErrEntityNotFound)
}

func TestStorage_ListEvents(t *testing.T) {
	s := New()
	ctx := context.Background()

	now := time.Now()
	userID := "user1"

	event1 := entity.Event{
		ID:        uuid.New().String(),
		Title:     "Event 1",
		TimeStart: now.Add(1 * time.Hour),
		TimeEnd:   now.Add(2 * time.Hour),
		UserId:    userID,
	}

	event2 := entity.Event{
		ID:        uuid.New().String(),
		Title:     "Event 2",
		TimeStart: now.Add(3 * time.Hour),
		TimeEnd:   now.Add(4 * time.Hour),
		UserId:    userID,
	}

	event3 := entity.Event{
		ID:        uuid.New().String(),
		Title:     "Event 3",
		TimeStart: now.Add(10 * time.Hour),
		TimeEnd:   now.Add(11 * time.Hour),
		UserId:    userID,
	}

	err := s.CreateEvent(ctx, event1)
	require.NoError(t, err)

	err = s.CreateEvent(ctx, event2)
	require.NoError(t, err)

	err = s.CreateEvent(ctx, event3)
	require.NoError(t, err)

	list, err := s.ListEvents(ctx, userID, now, now.Add(5*time.Hour))
	require.NoError(t, err)
	require.Len(t, list, 2)
}
