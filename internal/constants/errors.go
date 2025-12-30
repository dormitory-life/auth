package constants

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrConflict            = errors.New("conflict")
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("not found")
)
