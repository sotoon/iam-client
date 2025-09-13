package models

import (
	"errors"
)

var (
	ErrNotMatched          = errors.New("not matched")
	ErrForbidden           = errors.New("forbidden")
	ErrNotFound            = errors.New("not exists")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrBadRequest          = errors.New("bad request")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrCircuitBreakerOpen  = errors.New("circuit breaker is open")
	ErrInternalServerError = errors.New("internal server error")
)
