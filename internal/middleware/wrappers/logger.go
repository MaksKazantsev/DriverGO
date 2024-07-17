package wrappers

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/log"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func EmbedLogger(l log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.WithValue(c.Context(), utils.LoggerKey, l)
		c.SetUserContext(ctx)

		if logger := ctx.Value(utils.LoggerKey); logger == nil {
			l.Error("Logger not set in context", nil)
		} else {
			l.Info("Logger set in context", nil)
		}

		_ = c.Next()
		return nil
	}
}
