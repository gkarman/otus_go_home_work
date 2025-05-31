package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain"
)

type Storage struct {
	mu     sync.RWMutex
	events map[string]domain.Event
}

func New() *Storage {
	return &Storage{
		events: make(map[string]domain.Event),
	}
}

func (s *Storage) CreateEvent(ctx context.Context, event domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.events[event.ID]; exists {
		return domain.ErrEntityAlreadyExists
	}

	s.events[event.ID] = event
	return nil
}

func (s *Storage) UpdateEvent(ctx context.Context, event domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.events[event.ID]; !exists {
		return domain.ErrEntityNotFound
	}

	s.events[event.ID] = event
	return nil
}

func (s *Storage) DeleteEvent(ctx context.Context, eventID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.events[eventID]; !exists {
		return domain.ErrEntityNotFound
	}

	delete(s.events, eventID)
	return nil
}

func (s *Storage) ListEvents(ctx context.Context, userID string, from, to time.Time) ([]domain.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []domain.Event
	for _, event := range s.events {
		if event.UserId == userID && event.TimeStart.After(from) && event.TimeStart.Before(to) {
			result = append(result, event)
		}
	}

	return result, nil
}
