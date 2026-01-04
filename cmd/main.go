package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dormitory-life/auth/internal/config"
	"github.com/dormitory-life/auth/internal/database"
	"github.com/dormitory-life/auth/internal/grpc"
	"github.com/dormitory-life/auth/internal/logger"
	"github.com/dormitory-life/auth/internal/server"
	auth "github.com/dormitory-life/auth/internal/service"
)

func main() {
	configPath := os.Args[1]
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		panic(err)
	}

	log.Println("CONFIG: ", cfg)

	logger, err := logger.New(cfg)
	if err != nil {
		panic(err)
	}

	log.Println("Auth init db...")
	db, err := database.InitDb(cfg.Db)
	if err != nil {
		panic(err)
	}

	repository := database.New(db)

	authService := auth.New(auth.AuthServiceConfig{
		Repository: repository,
		JWTSecret:  cfg.JWT.Secret,
	})

	grpcServer := grpc.NewServer(grpc.GRPCServerConfig{
		AuthService: authService,
		Logger:      logger,
		Port:        cfg.GRPCServer.Port,
	})

	httpServer := server.New(server.ServerConfig{
		Config:      cfg.Server,
		AuthService: authService,
		Logger:      logger,
	})

	go func() {
		if err := grpcServer.Start(); err != nil {
			logger.Error("gRPC server failed", slog.String("error", err.Error()))
		}

		panic(err)
	}()

	go func() {
		if err := httpServer.Start(); err != nil {
			logger.Error("http server failed", slog.String("error", err.Error()))
		}

		panic(err)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
