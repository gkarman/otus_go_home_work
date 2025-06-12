package internalhttp

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/config"
)

type Server struct {
	cfg        config.ServerConf
	logger     logger.Logger
	app        Application
	httpServer *http.Server
}

type Application interface { // TODO
}

func New(cfg config.ServerConf, logger logger.Logger, app Application) *Server {
	return &Server{
		logger: logger,
		app:    app,
		cfg:    cfg,
	}
}

func (s *Server) Start(_ context.Context) error {
	mux := http.NewServeMux()

	mux.Handle("/hello", loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})))

	address := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	s.httpServer = &http.Server{
		Addr:              address,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	s.logger.Info("HTTP server starting at " + address)
	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
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
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		s.logger.Error("HTTP shutdown error: " + err.Error())
		return err
	}

	s.logger.Info("HTTP server stopped gracefully")
	return nil
}
