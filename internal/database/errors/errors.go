package dberrors

import "errors"

const (
	PGErrUniqueViolation = "23505"
)

var (
	ErrInternal   = errors.New("internal server error")
	ErrConflict   = errors.New("conflict")
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
)
