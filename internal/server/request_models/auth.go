package requestmodels

import (
	dbtypes "github.com/dormitory-life/auth/internal/database/types"
)

type (
	RegisterRequest struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		DormitoryId int    `json:"dormitory_id"`
	}

	RegisterResponse struct {
		UserId      string `json:"user_id"`
		DormitoryId int    `json:"dormitory_id"`
	}
)

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		UserId      string `json:"user_id"`
		DormitoryId int    `json:"dormitory_id"`
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

type ErrorResponse struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
}
