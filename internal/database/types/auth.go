package dbtypes

import "time"

type User struct {
	UserId      string
	Email       string
	Password    string
	DormitoryId string
	Role        string
	CreatedAt   time.Time
}

type (
	RegisterRequest struct {
		Email       string
		Password    string
		DormitoryId string
	}

	RegisterResponse struct {
		UserId      string
		DormitoryId string
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
		DormitoryId string
		Role        string
		CreatedAt   time.Time
	}
)

type (
	GetUserInfoByIdRequest struct {
		Id string
	}

	GetUserInfoByIdResponse struct {
		UserId      string
		DormitoryId string
		Role        string
	}
)
