package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dormitory-life/auth/internal/constants"
	dberrors "github.com/dormitory-life/auth/internal/database/errors"
	dbtypes "github.com/dormitory-life/auth/internal/database/types"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (c *Database) Register(
	ctx context.Context,
	request *dbtypes.RegisterRequest,
) (*dbtypes.RegisterResponse, error) {
	if request == nil {
		return nil, dberrors.ErrBadRequest
	}

	resp, err := c.register(ctx, c.db, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Database) register(
	ctx context.Context,
	driver Driver,
	request *dbtypes.RegisterRequest,
) (*dbtypes.RegisterResponse, error) {
	if request == nil {
		return nil, dberrors.ErrBadRequest
	}

	var (
		psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		usersTable = fmt.Sprintf("%s.%s", constants.SchemaName, constants.UsersTableName)
	)

	uid := uuid.NewString()

	queryBuilder := psql.Insert(usersTable).
		Columns(
			"id", "email", "password", "dormitory_id",
		).
		Values(
			uid, request.Email, request.Password, request.DormitoryId,
		)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: error building register query: %v", dberrors.ErrInternal, err)
	}
	_, err = driver.ExecContext(ctx, query, args...)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == dberrors.PGErrUniqueViolation {
			return nil, fmt.Errorf("%w: error user with same credentials exists: %v", dberrors.ErrConflict, err)
		}

		return nil, fmt.Errorf("%w: error executing register query: %v", dberrors.ErrInternal, err)
	}

	return &dbtypes.RegisterResponse{
		UserId:      uid,
		DormitoryId: request.DormitoryId,
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
		Select("id", "email", "password", "dormitory_id", "created_at").
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
		CreatedAt:   user.CreatedAt,
	}, nil
}
