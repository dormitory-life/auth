package main

import (
	"os"

	"github.com/dormitory-life/auth/internal/config"
	"github.com/dormitory-life/auth/internal/database"
	"github.com/dormitory-life/auth/internal/database/postgres"
	"github.com/dormitory-life/auth/internal/server"
	"github.com/dormitory-life/auth/internal/services/auth"
)

func main() {
	configPath := os.Args[1]
	config, err := config.ParseConfig(configPath)
	if err != nil {
		panic(err)
	}

	pgdb, err := postgres.New(*config)
	if err != nil {
		panic(err)
	}

	dbClient, err := database.New(pgdb)
	if err != nil {
		panic(err)
	}

	authSvc := auth.New(dbClient)

	// Здесь потом будет HTTP сервер
	s := server.New(*config, authSvc)
	panic(s.Start())
}
