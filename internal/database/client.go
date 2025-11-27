package database

import (
	"github.com/dormitory-life/auth/internal/database/postgres"
	_ "github.com/lib/pq"
)

type DbClient interface {
}

type Database struct {
	dbClient DbClient
}

func New(pgdb *postgres.PostgresDb) (*Database, error) {
	return &Database{
		dbClient: pgdb,
	}, nil
}
