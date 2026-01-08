package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dormitory-life/auth/internal/constants"
	dberrors "github.com/dormitory-life/auth/internal/database/errors"
	dbtypes "github.com/dormitory-life/auth/internal/database/types"
)

func (c *Database) GetUserById(
	ctx context.Context,
	request *dbtypes.GetUserInfoByIdRequest,
) (*dbtypes.GetUserInfoByIdResponse, error) {
	if request == nil {
		return nil, dberrors.ErrBadRequest
	}

	resp, err := c.getUserById(ctx, c.db, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Database) getUserById(
	ctx context.Context,
	driver Driver,
	request *dbtypes.GetUserInfoByIdRequest,
) (*dbtypes.GetUserInfoByIdResponse, error) {
	if request == nil {
		return nil, dberrors.ErrBadRequest
	}

	var (
		psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		usersTable = fmt.Sprintf("%s.%s", constants.SchemaName, constants.UsersTableName)
	)

	queryBuilder := psql.
		Select("id", "dormitory_id", "role").
		From(usersTable).
		Where(squirrel.Eq{"id": request.Id}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: error building get user by id query: %v", dberrors.ErrInternal, err)
	}

	var user dbtypes.User
	err = driver.QueryRowContext(ctx, query, args...).Scan(
		&user.UserId,
		&user.DormitoryId,
		&user.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: user not found", dberrors.ErrNotFound)
		}

		return nil, fmt.Errorf("%w: error executing get user by id query: %v", dberrors.ErrInternal, err)
	}

	return &dbtypes.GetUserInfoByIdResponse{
		UserId:      user.UserId,
		DormitoryId: user.DormitoryId,
		Role:        user.Role,
	}, nil
}

func (c *Database) GetUserByEmail(
	ctx context.Context,
	request *dbtypes.GetUserByEmailRequest,
) (*dbtypes.GetUserResponse, error) {
	if request == nil {
		return nil, dberrors.ErrBadRequest
	}

	resp, err := c.getUserByEmail(ctx, c.db, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Database) getUserByEmail(
	ctx context.Context,
	driver Driver,
	request *dbtypes.GetUserByEmailRequest,
) (*dbtypes.GetUserResponse, error) {
	if request == nil {
		return nil, dberrors.ErrBadRequest
	}

	var (
		psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		usersTable = fmt.Sprintf("%s.%s", constants.SchemaName, constants.UsersTableName)
	)

	queryBuilder := psql.
		Select("id", "email", "password", "dormitory_id", "role", "created_at").
		From(usersTable).
		Where(squirrel.Eq{"email": request.Email}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: error building get user query: %v", dberrors.ErrInternal, err)
	}

	var user dbtypes.User
	err = driver.QueryRowContext(ctx, query, args...).Scan(
		&user.UserId,
		&user.Email,
		&user.Password,
		&user.DormitoryId,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: user not found", dberrors.ErrNotFound)
		}

		return nil, fmt.Errorf("%w: error executing get user query: %v", dberrors.ErrInternal, err)
	}

	return &dbtypes.GetUserResponse{
		UserId:      user.UserId,
		Email:       user.Email,
		Password:    user.Password,
		DormitoryId: user.DormitoryId,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
	}, nil
}
