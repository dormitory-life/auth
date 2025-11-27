package server

import (
	"fmt"
	"net/http"

	config "github.com/dormitory-life/auth/internal/config"
	"github.com/dormitory-life/auth/internal/services/auth"
)

type Server struct {
	server http.Server
	auth   *auth.AuthSvc
}

func New(cfg config.Config, authSvc *auth.AuthSvc) *Server {
	s := new(Server)
	s.server.Addr = fmt.Sprintf(":%d", cfg.Server.Port)
	s.server.Handler = s.setupRouter()
	s.auth = authSvc
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
