package main

import (
	"log"
	"os"

	"github.com/dormitory-life/auth/internal/config"
	"github.com/dormitory-life/auth/internal/database"
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

	db, err := database.InitDb(cfg.Db)
	if err != nil {
		panic(err)
	}

	repository := database.New(db)

	authService := auth.New(auth.AuthServiceConfig{
		Repository: repository,
		JWTSecret:  cfg.JWT.Secret,
	})

	s := server.New(server.ServerConfig{
		Config:      cfg.Server,
		AuthService: authService,
		Logger:      logger,
	})

	panic(s.Start())
}
