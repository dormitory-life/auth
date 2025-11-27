package auth

import (
	db "github.com/dormitory-life/auth/internal/database"
)

type AuthSvc struct {
	DbClient db.DbClient
}

func New(dbClient db.DbClient) *AuthSvc {
	return &AuthSvc{
		DbClient: dbClient,
	}
}
