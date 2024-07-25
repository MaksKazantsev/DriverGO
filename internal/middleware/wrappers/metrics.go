package wrappers

import (
	"github.com/MaksKazantsev/DriverGO/internal/metrics"
	"github.com/gofiber/fiber/v2"
	"time"
)

func WithRequestLatency(m metrics.Metrics) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		_ = ctx.Next()

		duration := time.Since(start).Seconds()

		m.Latency(duration, ctx.Path())

		return ctx.Next()
	}
}
