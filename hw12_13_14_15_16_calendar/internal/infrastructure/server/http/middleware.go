package internalhttp

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
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

		file, err := os.OpenFile("logs/http.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			fmt.Println("could not open access.log:", err)
			return
		}
		defer file.Close()

		_, _ = file.WriteString(logMessage + "\n")
	})
}
