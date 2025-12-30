package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	rmodel "github.com/dormitory-life/auth/internal/server/request_models"
	auth "github.com/dormitory-life/auth/internal/service"

	"github.com/dormitory-life/auth/internal/constants"
)

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (s *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
	const handlerName = "registerHandler"

	var req rmodel.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		s.logger.Error("error decoding request",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)

		return
	}

	s.logger.Debug("request", slog.Any("req", req))

	resp, err := s.authService.Register(r.Context(), &req)
	if err != nil {
		s.handleError(w, err)
		s.logger.Error("error",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)

		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		s.logger.Error("error encoding response",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)
	}
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	const handlerName = "loginHandler"

	var req rmodel.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		s.logger.Error("error decoding request",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)

		return
	}

	s.logger.Debug("request", slog.Any("req", req))

	resp, err := s.authService.Login(r.Context(), &req)
	if err != nil {
		s.handleError(w, err)
		s.logger.Error("error",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)

		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		s.logger.Error("error encoding response",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)
	}
}

func writeErrorResponse(w http.ResponseWriter, err error, code int, details ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := rmodel.ErrorResponse{
		Error:   err.Error(),
		Details: details,
	}

	_ = json.NewEncoder(w).Encode(response)
}

func (s *Server) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, auth.ErrBadRequest):
		writeErrorResponse(w, constants.ErrBadRequest, http.StatusBadRequest, err.Error())
	case errors.Is(err, auth.ErrNotFound):
		writeErrorResponse(w, constants.ErrNotFound, http.StatusNotFound)
	case errors.Is(err, auth.ErrConflict):
		writeErrorResponse(w, constants.ErrConflict, http.StatusConflict, err.Error())
	case errors.Is(err, auth.ErrInternal):
		writeErrorResponse(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	default:
		s.logger.Error("Unhandled auth error", slog.String("error", err.Error()))
		writeErrorResponse(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}
}
