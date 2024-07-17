package utils

import (
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.NewError(errors.ERR_INTERNAL, "failed to hash password: "+err.Error())
	}
	return string(b), nil
}

func CompareHash(savedPassword, reqPassword string) error {
	fmt.Println(savedPassword, reqPassword)
	if err := bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(reqPassword)); err != nil {
		return errors.NewError(errors.ERR_BAD_REQUEST, "wrong password provided")
	}
	return nil
}
