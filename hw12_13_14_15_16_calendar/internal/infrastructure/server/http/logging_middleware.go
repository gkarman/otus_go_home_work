package internalhttp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler, logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		latency := time.Since(start)
		clientIP := r.RemoteAddr
		userAgent := r.UserAgent()

		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)

		logMessage := fmt.Sprintf(
			`%s [%s] %s %s %s %d %s "%s"`,
			clientIP,
			start.Format("02/Jan/2001:01:01:01 -0700"),
			r.Method,
			r.URL.RequestURI(),
			r.Proto,
			rw.statusCode,
			latency,
			userAgent,
		)

		logger.LogToFile(logMessage)
	})
}
