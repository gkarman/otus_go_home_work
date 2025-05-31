package sqlstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func New(dsn string) (*Storage, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Storage) CreateEvent(ctx context.Context, event domain.Event) error {
	query := `
		INSERT INTO events (id, title, time_start, time_end, description, user_id, notify_before)
		VALUES (:id, :title, :time_start, :time_end, :description, :user_id, :notify_before)
	`
	_, err := s.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":            event.ID,
		"title":         event.Title,
		"time_start":    event.TimeStart,
		"time_end":      event.TimeEnd,
		"description":   event.Description,
		"user_id":       event.UserId,
		"notify_before": int64(event.NotifyBefore.Seconds()),
	})
	if err != nil {
		return fmt.Errorf("create event: %w", err)
	}
	return nil
}

func (s *Storage) UpdateEvent(ctx context.Context, event domain.Event) error {
	query := `
		UPDATE events
		SET title=:title, time_start=:time_start, time_end=:time_end,
		    description=:description, user_id=:user_id, notify_before=:notify_before
		WHERE id=:id
	`
	res, err := s.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":            event.ID,
		"title":         event.Title,
		"time_start":    event.TimeStart,
		"time_end":      event.TimeEnd,
		"description":   event.Description,
		"user_id":       event.UserId,
		"notify_before": int64(event.NotifyBefore.Seconds()),
	})
	if err != nil {
		return fmt.Errorf("update event: %w", err)
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return domain.ErrEntityNotFound
	}
	return nil
}

func (s *Storage) DeleteEvent(ctx context.Context, eventID string) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM events WHERE id = $1`, eventID)
	if err != nil {
		return fmt.Errorf("delete event: %w", err)
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return domain.ErrEntityNotFound
	}
	return nil
}

func (s *Storage) ListEvents(ctx context.Context, userID string, from, to time.Time) ([]domain.Event, error) {
	var events []domain.Event

	query := `
		SELECT id, title, time_start, time_end, description, user_id, notify_before
		FROM events
		WHERE user_id = $1 AND time_start >= $2 AND time_start <= $3
	`

	err := s.db.SelectContext(ctx, &events, query, userID, from, to)
	if err != nil {
		return nil, fmt.Errorf("list events: %w", err)
	}

	for i := range events {
		events[i].NotifyBefore = time.Duration(events[i].NotifyBefore) * time.Second
	}

	return events, nil
}
