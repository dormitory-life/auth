package main

import (
	"os"

	"github.com/dormitory-life/auth/internal/config"
	"github.com/dormitory-life/auth/internal/database/postgres"
	"github.com/dormitory-life/auth/internal/repository"
	"github.com/dormitory-life/auth/internal/server"
	"github.com/dormitory-life/auth/internal/service"
)

func main() {
	configPath := os.Args[1]
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		panic(err)
	}

	// 1. Подключаемся к БД
	db, err := postgres.New(*cfg)
	if err != nil {
		panic(err)
	}

	// 2. Создаем репозиторий
	userRepo := repository.NewUserRepository(db.DB)

	// 3. Создаем сервис
	authService := services.NewAuthService(userRepo, "your-jwt-secret")

	// 4. Запускаем сервер
	s := server.New(*cfg, authService)
	panic(s.Start())
}