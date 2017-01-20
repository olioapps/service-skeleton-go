package api

import (
	log "github.com/Sirupsen/logrus"
)

const (
	BAD_REQUEST_EXCEPTION  = 400
	UNAUTHORIZED_EXCEPTION = 401
	FORBIDDEN_EXCEPTION    = 403
	NOT_FOUND_EXCEPTION    = 404
	CONFLICT_EXCEPTION     = 409
	RUNTIME_EXCEPTION      = 500
)

type Exception struct {
	ErrorCode int
	Err       string
}

func (e *Exception) Error() string {
	return e.Err
}

func NewException(errorCode int, err string) *Exception {
	obj := Exception{ErrorCode: errorCode, Err: err}

	return &obj
}

func NewBadRequestException(err string) *Exception {
	log.Error(err)
	return NewException(BAD_REQUEST_EXCEPTION, err)
}

func NewUnauthorizedException(err string) *Exception {
	log.Error(err)
	return NewException(UNAUTHORIZED_EXCEPTION, err)
}

func NewForbiddenException(err string) *Exception {
	log.Error(err)
	return NewException(FORBIDDEN_EXCEPTION, err)
}

func NewConflictException(err string) *Exception {
	log.Error(err)
	return NewException(CONFLICT_EXCEPTION, err)
}

func NewNotFoundException(err string) *Exception {
	log.Error(err)
	return NewException(NOT_FOUND_EXCEPTION, err)
}

func NewRuntimeException(err string) *Exception {
	log.Error(err)
	return NewException(RUNTIME_EXCEPTION, err)
}
