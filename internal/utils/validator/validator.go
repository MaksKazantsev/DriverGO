package validator

import (
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"regexp"
)

type Validator interface {
	ValidateRegistration(data models.RegisterReq) error
	ValidateLogin(data models.LoginReq) error
}

func NewValidator() Validator {
	return &validator{
		regExpEmail: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	}
}

type validator struct {
	regExpEmail *regexp.Regexp
}

func (v validator) ValidateLogin(data models.LoginReq) error {
	if ok := v.regExpEmail.MatchString(data.Email); !ok {
		return errors.NewError(errors.ERR_BAD_REQUEST, "invalid email")
	}
	if len(data.Password) < 5 || len(data.Password) > 30 {
		return errors.NewError(errors.ERR_BAD_REQUEST, "password does not match all options")
	}
	return nil
}

func (v validator) ValidateRegistration(data models.RegisterReq) error {
	if ok := v.regExpEmail.MatchString(data.Email); !ok {
		return errors.NewError(errors.ERR_BAD_REQUEST, "invalid email")
	}

	if len(data.Password) < 5 || len(data.Password) > 30 {
		return errors.NewError(errors.ERR_BAD_REQUEST, "password does not match all options")
	}

	if len(data.Username) < 3 || len(data.Username) > 25 {
		return errors.NewError(errors.ERR_BAD_REQUEST, "username does not match all options")
	}
	return nil
}
