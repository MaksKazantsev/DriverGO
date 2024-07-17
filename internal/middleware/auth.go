package middleware

import (
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"time"
)

func CheckAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		values := ctx.Get("Authorization")
		valuesArr := strings.Split(values, " ")

		if len(valuesArr) != 2 {
			_ = ctx.Status(http.StatusMethodNotAllowed).SendString("token is not provided")
			return nil
		}

		if valuesArr[0] != "Bearer" {
			_ = ctx.Status(http.StatusBadRequest).SendString("wrong request signature")
			return nil
		}

		claims, err := utils.ParseToken(valuesArr[1])
		if err != nil {
			_ = ctx.Status(http.StatusMethodNotAllowed).SendString(err.Error())
		}

		_, ok := claims["id"].(string)
		if !ok {
			_ = ctx.Status(http.StatusMethodNotAllowed).SendString("invalid token")
			return nil
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			_ = ctx.Status(http.StatusMethodNotAllowed).SendString("invalid token")
			return nil
		}

		if time.Now().After(time.Unix(int64(exp), 0)) {
			_ = ctx.Status(http.StatusMethodNotAllowed).SendString("token is expired")
			return nil
		}
		_ = ctx.Next()
		return nil
	}
}
