package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/dormitory-life/auth/internal/database"
	dberrors "github.com/dormitory-life/auth/internal/database/errors"
	dbtypes "github.com/dormitory-life/auth/internal/database/types"
	rmodel "github.com/dormitory-life/auth/internal/server/request_models"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceConfig struct {
	Repository database.Repository
	JWTSecret  string
}
type AuthService struct {
	repository database.Repository
	jwtSecret  string
}

type AuthServiceClient interface {
	Register(ctx context.Context, request *rmodel.RegisterRequest) (*rmodel.RegisterResponse, error)
	Login(ctx context.Context, request *rmodel.LoginRequest) (*rmodel.LoginResponse, error)
	RefreshTokens(ctx context.Context, request *rmodel.RefreshTokensRequest) (*rmodel.RefreshTokensResponse, error)

	GetUserInfoById(ctx context.Context, request *rmodel.GetUserByIdRequest) (*rmodel.GetUserByIdResponse, error)
}

func New(cfg AuthServiceConfig) AuthServiceClient {
	return &AuthService{
		repository: cfg.Repository,
		jwtSecret:  cfg.JWTSecret,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	request *rmodel.RegisterRequest,
) (*rmodel.RegisterResponse, error) {
	if request == nil {
		return nil, ErrBadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	resp, err := s.repository.Register(ctx, &dbtypes.RegisterRequest{
		Email:       request.Email,
		Password:    string(hashedPassword),
		DormitoryId: request.DormitoryId,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: error register user: %v", s.handleDBError(err), err)
	}

	result := new(rmodel.RegisterResponse).From(resp)

	accessToken, refreshToken, err := s.generateJWTTokens(ctx, result.UserId, result.DormitoryId)
	if err != nil {
		return nil, fmt.Errorf("%w: error register user: %v", s.handleDBError(err), err)
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken

	return result, nil
}

func (s *AuthService) Login(
	ctx context.Context,
	request *rmodel.LoginRequest,
) (*rmodel.LoginResponse, error) {
	if request == nil {
		return nil, ErrBadRequest
	}

	resp, err := s.repository.GetUserByEmail(ctx, &dbtypes.GetUserByEmailRequest{
		Email: request.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: error getting user while login: %v", s.handleDBError(err), err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(request.Password))
	if err != nil {
		return nil, fmt.Errorf("%w: incorrect password", ErrUnauthorized)
	}

	result := &rmodel.LoginResponse{
		UserId:      resp.UserId,
		DormitoryId: resp.DormitoryId,
	}

	accessToken, refreshToken, err := s.generateJWTTokens(ctx, result.UserId, result.DormitoryId)
	if err != nil {
		return nil, fmt.Errorf("%w: error register user: %v", s.handleDBError(err), err)
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken

	return result, nil
}

func (s *AuthService) RefreshTokens(
	ctx context.Context,
	request *rmodel.RefreshTokensRequest,
) (*rmodel.RefreshTokensResponse, error) {
	if request == nil {
		return nil, ErrBadRequest
	}

	accessToken, refreshToken, err := s.generateJWTTokens(ctx, request.UserId, request.DormitoryId)
	if err != nil {
		return nil, fmt.Errorf("%w: error refreshing tokens: %v", ErrInternal, err)
	}

	return &rmodel.RefreshTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) handleDBError(err error) error {
	switch {
	case errors.Is(err, dberrors.ErrBadRequest):
		return ErrBadRequest
	case errors.Is(err, dberrors.ErrNotFound):
		return ErrNotFound
	case errors.Is(err, dberrors.ErrInternal):
		return ErrInternal
	case errors.Is(err, dberrors.ErrConflict):
		return ErrConflict
	default:
		return ErrInternal
	}
}
