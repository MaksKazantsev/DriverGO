package utils

import (
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

type TokenType string

const (
	ACCESS  TokenType = "access"
	REFRESH TokenType = "refresh"
)

var KEY = os.Getenv("TOKEN_SECRET_KEY")

type TokenData struct {
	ID   string
	Role string
}

func ParseDuration(duration string, multiplier time.Duration) time.Duration {
	d, err := time.ParseDuration(duration)
	if err != nil {
		log.Println(err)
	}
	return d * multiplier
}

func NewToken(t TokenType, data TokenData) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	cl := token.Claims.(jwt.MapClaims)
	switch t {
	case ACCESS:
		cl["id"] = data.ID
		cl["role"] = data.Role
		cl["exp"] = time.Now().Add(time.Minute * 60).Unix()
	case REFRESH:
		cl["id"] = data.ID
		cl["role"] = data.Role
		cl["exp"] = time.Now().Add(time.Minute * 60 * 24 * 30).Unix()
	}

	stringToken, err := token.SignedString([]byte("TOKEN_SECRET_KEY"))
	if err != nil {
		return "", errors.NewError(errors.ERR_INTERNAL, "failed to sign token: "+err.Error())
	}
	return stringToken, nil
}

func ParseToken(token string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("TOKEN_SECRET_KEY"), nil
	})
	if err != nil {
		return nil, errors.NewError(errors.ERR_NOT_ALLOWED, "invalid token")
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if claims.Valid() != nil || !ok {
		return nil, errors.NewError(errors.ERR_NOT_ALLOWED, "invalid token")
	}
	return claims, nil
}
