package internalhttp

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/config"
)

const timeout = 5 * time.Second

type Server struct {
	cfg        config.ServerConf
	logger     logger.Logger
	app        application.Calendar
	httpServer *http.Server
}

func New(cfg config.ServerConf, logger logger.Logger, app application.Calendar) *Server {
	return &Server{
		logger: logger,
		cfg:    cfg,
		app:    app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	mux := http.NewServeMux()

	mux.Handle("/hello", loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	}), s.logger))

	mux.Handle("/create", loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {

		requestDto := requestdto.CreateEvent{
			UserID:       "11111111-1111-1111-1111-111111111111",
			Title:        "Meeting with team",
			Description:  "Discuss project updates",
			TimeStart:    time.Date(2025, 6, 10, 14, 0, 0, 0, time.UTC),
			TimeEnd:      time.Date(2025, 6, 10, 15, 0, 0, 0, time.UTC),
			NotifyBefore: 15 * time.Minute,
		}

		err := s.app.CreateEvent(ctx, requestDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
		}

	}), s.logger))

	address := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	s.httpServer = &http.Server{
		Addr:              address,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	s.logger.Info("HTTP server starting at " + address)
	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("HTTP server failed: " + err.Error())
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}
	s.logger.Info("Shutting down HTTP server...")

	// Graceful shutdown с таймаутом
	shutdownCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		s.logger.Error("HTTP shutdown error: " + err.Error())
		return err
	}

	s.logger.Info("HTTP server stopped gracefully")
	return nil
}
