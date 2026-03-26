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

// @Summary Проверка доступности auth-сервиса
// @Description Возвращает pong, если сервис авторизации работает
// @Tags auth
// @Produce json
// @Success 200 {string} string "pong"
// @Router /auth/ping [get]
func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// @Summary Регистрация нового пользователя
// @Description Создает нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param request body rmodel.RegisterRequest true "Данные пользователя для регистрации"
// @Success 201 {object} rmodel.RegisterResponse "Пользователь зарегистрирован"
// @Failure 400 {object} rmodel.ErrorResponse "Неверные данные / параметры запроса"
// @Failure 409 {object} rmodel.ErrorResponse "Пользователь с такими данными уже существует"
// @Failure 500 {object} rmodel.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/register [post]
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

	s.logger.Debug(handlerName, slog.Any("req", req))

	resp, err := s.authService.Register(r.Context(), &req)
	if err != nil {
		s.handleError(w, err)
		s.logger.Error("error",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		s.logger.Error("error encoding response",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)
	}
}

// @Summary Вход в систему
// @Description Залогинивает пользователя и выдает токены
// @Tags auth
// @Accept json
// @Produce json
// @Param request body rmodel.LoginRequest true "Данные пользователя для входа"
// @Success 200 {object} rmodel.LoginResponse "Пользователь авторизован"
// @Failure 400 {object} rmodel.ErrorResponse "Неверные данные / параметры запроса"
// @Failure 401 {object} rmodel.ErrorResponse "Неверные данные для входа"
// @Failure 500 {object} rmodel.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/login [post]
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

	s.logger.Debug(handlerName, slog.Any("req", req))

	resp, err := s.authService.Login(r.Context(), &req)
	if err != nil {
		s.handleError(w, err)
		s.logger.Error("error",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		writeErrorResponse(w, err, http.StatusInternalServerError)
		s.logger.Error("error encoding response",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)
	}
}

// @Summary Обновление токенов
// @Description Выдает пользователю пару новых токенов
// @Tags auth
// @Accept json
// @Produce json
// @Param request body rmodel.RefreshTokensRequest true "Пара старых токенов"
// @Success 200 {object} rmodel.RefreshTokensResponse "Новые токены выданы"
// @Failure 400 {object} rmodel.ErrorResponse "Неверные данные / параметры запроса"
// @Failure 500 {object} rmodel.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/refresh [post]
func (s *Server) refreshHandler(w http.ResponseWriter, r *http.Request) {
	const handlerName = "refreshHandler"

	userId := r.Header.Get("X-User-ID")
	dormitoryId := r.Header.Get("X-Dormitory-ID")

	if userId == "" || dormitoryId == "" {
		writeErrorResponse(w, constants.ErrBadRequest, http.StatusBadRequest, "Missing user data")
		return
	}

	var req rmodel.RefreshTokensRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, err, http.StatusBadRequest)
		s.logger.Error("error decoding request",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)

		return
	}
	req.UserId = userId
	req.DormitoryId = dormitoryId

	s.logger.Debug(handlerName, slog.Any("req", req))

	resp, err := s.authService.RefreshTokens(r.Context(), &req)
	if err != nil {
		s.handleError(w, err)
		s.logger.Error("error",
			slog.String("error", err.Error()),
			slog.String("handler", handlerName),
		)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

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
	case errors.Is(err, auth.ErrUnauthorized):
		writeErrorResponse(w, constants.ErrUnauthorized, http.StatusUnauthorized)
	case errors.Is(err, auth.ErrInternal):
		writeErrorResponse(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	default:
		s.logger.Error("Unhandled auth error", slog.String("error", err.Error()))
		writeErrorResponse(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}
}
