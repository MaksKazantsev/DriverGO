package errors

import (
	"context"
	"errors"
	"github.com/MaksKazantsev/DriverGO/internal/log"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
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
		utils.ExtractLogger(ctx).Error(err.Error(), &log.Data{})
		return http.StatusInternalServerError, "unknown error: " + err.Error()
	}

	switch e.Code {
	case ERR_INTERNAL:
		utils.ExtractLogger(ctx).Error(err.Error(), &log.Data{})
		return http.StatusInternalServerError, "internal error: " + err.Error()
	case ERR_NOT_ALLOWED:
		return http.StatusMethodNotAllowed, e.Message
	case ERR_BAD_REQUEST:
		return http.StatusBadRequest, e.Message
	case ERR_NOT_FOUND:
		return http.StatusNotFound, e.Message
	default:
		utils.ExtractLogger(ctx).Error(err.Error(), &log.Data{})
		return http.StatusInternalServerError, "unknown error: " + err.Error()
	}

}
