package server

import (
	"log/slog"
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("request started",
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Any("date", time.Now().UTC()))

		next.ServeHTTP(w, r)

		s.logger.Info("request completed",
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Any("date", time.Now().UTC()))
	})
}
