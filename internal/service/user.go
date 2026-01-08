package auth

import (
	"context"
	"fmt"

	dbtypes "github.com/dormitory-life/auth/internal/database/types"
	rmodel "github.com/dormitory-life/auth/internal/server/request_models"
)

func (s *AuthService) GetUserInfoById(
	ctx context.Context,
	request *rmodel.GetUserByIdRequest,
) (*rmodel.GetUserByIdResponse, error) {
	if request == nil {
		return nil, ErrBadRequest
	}

	resp, err := s.repository.GetUserById(ctx, &dbtypes.GetUserInfoByIdRequest{
		Id: request.UserId,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: error getting user by id: %v", s.handleDBError(err), err)
	}

	result := &rmodel.GetUserByIdResponse{
		Info: &rmodel.UserInfoById{
			UserId:      resp.UserId,
			DormitoryId: resp.DormitoryId,
			Role:        resp.Role,
		},
	}

	return result, nil
}
