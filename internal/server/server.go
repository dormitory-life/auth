package server

import (
	"fmt"
	"net/http"

	"github.com/dormitory-life/auth/internal/config"
	"github.com/dormitory-life/auth/internal/service"
)

type ServerConfig struct {
	Config  config.ServerConfig
	Service service.Service
}

type Server struct {
	server  http.Server
	service service.Service
}

func New(cfg ServerConfig) *Server {
	s := new(Server)
	s.server.Addr = fmt.Sprintf("%d", cfg.Config.Port)
	s.server.Handler = s.setupRouter()
	s.service = cfg.Service

	return s
}

func (s *Server) setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", s.pingHandler)

	return mux
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
