package errors

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func ErrorDBWrapper(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewError(ERR_NOT_FOUND, "entity not found: "+err.Error())
	}
	return NewError(ERR_INTERNAL, "internal db error: "+err.Error())
}

func ErrorRepoWrapper(err error) error {
	return fmt.Errorf("repo error: %w", err)
}
