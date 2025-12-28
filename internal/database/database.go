package database

import (
	"database/sql"
	"log"

	"github.com/dormitory-life/auth/internal/config"
	"github.com/dormitory-life/utils/migrator"
)

type database struct {
	db *sql.DB
}

type Repository interface {
	//funcs
}

func New(db *sql.DB) Repository {
	return &database{
		db: db,
	}
}

func InitDb(cfg config.DataBaseConfig) (*sql.DB, error) {
	connStr := cfg.GetConnectionString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if err := migrator.MigrateDB(connStr, cfg.MigrationsPath); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migrations completed successfully!")
	log.Println("Auth service is ready")

	return db, nil
}
