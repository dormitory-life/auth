package requestmodels

import (
	"time"

	dbtypes "github.com/dormitory-life/auth/internal/database/types"
)

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
		Email       string `json:"email"`
		Password    string `json:"password"`
		DormitoryId string `json:"dormitory_id"`
	}

	RegisterResponse struct {
		UserId       string `json:"user_id"`
		DormitoryId  string `json:"dormitory_id"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		UserId       string `json:"user_id"`
		DormitoryId  string `json:"dormitory_id"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func (*RegisterResponse) From(msg *dbtypes.RegisterResponse) *RegisterResponse {
	if msg == nil {
		return nil
	}

	return &RegisterResponse{
		UserId:      msg.UserId,
		DormitoryId: msg.DormitoryId,
	}
}

type UserInfoById struct {
	UserId      string
	DormitoryId string
	Role        string
}

type (
	GetUserByIdRequest struct {
		UserId string `json:"user_id"`
	}

	GetUserByIdResponse struct {
		Info *UserInfoById
	}
)
