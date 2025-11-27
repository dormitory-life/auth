package postgres

import (
	"database/sql"
	"log"

	"github.com/dormitory-life/auth/internal/config"
	"github.com/dormitory-life/utils/migrator"
)

type PostgresDb struct {
	db *sql.DB
}

func New(cfg config.Config) (*PostgresDb, error) {
	connStr := cfg.Db.GetConnectionString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if err := migrator.MigrateDB(connStr, cfg.Db.MigrationsPath); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migrations completed successfully!")
	log.Println("Auth service is ready")

	return &PostgresDb{
		db: db,
	}, nil
}
