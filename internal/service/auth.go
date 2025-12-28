package service

import (
	"github.com/dormitory-life/auth/internal/database"
)

type authService struct {
	repository database.Repository
}

type AuthService interface {
}

func New(repository database.Repository) AuthService {
	return &authService{
		repository: repository,
	}
}
