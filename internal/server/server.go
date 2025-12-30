package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dormitory-life/auth/internal/config"
	auth "github.com/dormitory-life/auth/internal/service"
)

type ServerConfig struct {
	Config      config.ServerConfig
	AuthService auth.AuthServiceClient
	Logger      *slog.Logger
}

type Server struct {
	server      http.Server
	authService auth.AuthServiceClient
	logger      *slog.Logger
}

func New(cfg ServerConfig) *Server {
	s := new(Server)
	s.server.Addr = fmt.Sprintf(":%d", cfg.Config.Port)
	s.server.Handler = s.setupRouter()
	s.authService = cfg.AuthService
	s.logger = cfg.Logger

	return s
}

func (s *Server) setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /auth/ping", s.pingHandler)
	mux.HandleFunc("POST /auth/register", s.registerHandler)
	mux.HandleFunc("POST /auth/login", s.loginHandler)
	mux.HandleFunc("POST /auth/refresh", s.refreshHandler)

	return mux
}

func (s *Server) Start() error {
	s.logger.Debug("server started", slog.String("address", s.server.Addr))
	return s.server.ListenAndServe()
}
