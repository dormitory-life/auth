package dbtypes

import "time"

type User struct {
	UserId      string
	Email       string
	Password    string
	DormitoryId int
	CreatedAt   time.Time
}

type (
	RegisterRequest struct {
		Email       string
		Password    string
		DormitoryId int
	}

	RegisterResponse struct {
		UserId      string
		DormitoryId int
	}
)

type (
	GetUserByEmailRequest struct {
		Email string
	}

	GetUserResponse struct {
		UserId      string
		Email       string
		Password    string
		DormitoryId int
		CreatedAt   time.Time
	}
)
