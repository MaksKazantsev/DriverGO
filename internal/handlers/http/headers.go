package http

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func extractAuthHeader(c *fiber.Ctx) string {
	values := c.Get("Authorization")
	valuesArr := strings.Split(values, " ")
	if len(valuesArr) != 2 {
		return ""
	}
	return valuesArr[1]
}
