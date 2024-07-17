package wrappers

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func WithIdempotencyKey() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		ctx = context.WithValue(ctx, utils.IdempotencyKey, uuid.NewString())
		c.SetUserContext(ctx)

		_ = c.Next()
		return nil
	}
}
