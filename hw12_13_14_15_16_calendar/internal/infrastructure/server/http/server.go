package internalhttp

import (
	"context"
	"encoding/json"
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

func (s *Server) Start(_ context.Context) error {
	mux := s.registerRoutes()
	wrappedMux := loggingMiddleware(mux, s.logger)

	address := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	s.httpServer = &http.Server{
		Addr:              address,
		Handler:           wrappedMux,
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

func (s *Server) registerRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", s.handleCreate())
	mux.HandleFunc("/delete", s.handleDelete())
	mux.HandleFunc("/update", s.handleUpdate())
	mux.HandleFunc("/events_day", s.handleEventsDay())
	mux.HandleFunc("/events_week", s.handleEventsWeek())
	mux.HandleFunc("/events_month", s.handleEventsMonth())

	return mux
}

func (s *Server) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestDto requestdto.CreateEvent
		if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
			http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		response, err := s.app.CreateEvent(ctx, requestDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.ID))
	}
}

//nolint:dupl // handleUpdate похож на handleDelete, но это ожидаемо
func (s *Server) handleDelete() http.HandlerFunc {
	type response struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestDto requestdto.DeleteEvent
		if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
			http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		err := s.app.DeleteEvent(ctx, requestDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{
			Status:  true,
			Message: "event deleted",
		})
	}
}

//nolint:dupl // handleUpdate похож на handleDelete, но это ожидаемо
func (s *Server) handleUpdate() http.HandlerFunc {
	type response struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestDto requestdto.UpdateEvent
		if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
			http.Error(w, "invalid update request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		err := s.app.UpdateEvent(ctx, requestDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{
			Status:  true,
			Message: "event updated",
		})
	}
}

//nolint:dupl
func (s *Server) handleEventsDay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestDto requestdto.EventsOnDate
		if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
			http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		resp, err := s.app.EventsDay(ctx, requestDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}

//nolint:dupl
func (s *Server) handleEventsWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestDto requestdto.EventsOnDate
		if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
			http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		resp, err := s.app.EventsWeek(ctx, requestDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}

//nolint:dupl
func (s *Server) handleEventsMonth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestDto requestdto.EventsOnDate
		if err := json.NewDecoder(r.Body).Decode(&requestDto); err != nil {
			http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		resp, err := s.app.EventsMonth(ctx, requestDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}
