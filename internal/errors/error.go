package errors

import (
	"context"
	"errors"
	"net/http"
)

type ErrorCode int

const (
	ERR_INTERNAL ErrorCode = iota
	ERR_NOT_FOUND
	ERR_NOT_ALLOWED
	ERR_BAD_REQUEST
)

type Error struct {
	Code    ErrorCode
	Message string
}

func IntToErrorCode(code int) ErrorCode {
	return ErrorCode(code)
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code ErrorCode, msg string) error {
	return &Error{Code: code, Message: msg}
}

func FromError(err error, ctx context.Context) (int, string) {
	var e *Error

	if !errors.As(err, &e) {
		return http.StatusInternalServerError, "unknown error: " + err.Error()
	}

	switch e.Code {
	case ERR_INTERNAL:
		return http.StatusInternalServerError, "internal error: " + err.Error()
	case ERR_NOT_ALLOWED:
		return http.StatusMethodNotAllowed, e.Message
	case ERR_BAD_REQUEST:
		return http.StatusBadRequest, e.Message
	case ERR_NOT_FOUND:
		return http.StatusNotFound, e.Message
	default:
		return http.StatusInternalServerError, "unknown error: " + err.Error()
	}

}
