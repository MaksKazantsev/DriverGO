package errors

import (
	"fmt"
	"strings"
)

func ErrorDBWrapper(err error) error {
	if strings.Contains(err.Error(), "not found") {
		return NewError(ERR_NOT_FOUND, "entity not found")
	}
	if strings.Contains(err.Error(), "duplicate key") {
		return NewError(ERR_BAD_REQUEST, "entity already exists")
	}
	return NewError(ERR_INTERNAL, "internal db error: "+err.Error())
}

func ErrorRepoWrapper(err error) error {
	return fmt.Errorf("repo error: %w", err)
}
