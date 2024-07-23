package utils

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/log"
)

func ExtractLogger(ctx context.Context) log.Logger {
	l, ok := ctx.Value(LoggerKey).(log.Logger)
	if !ok {
		panic("can not get logger from context")
	}
	return l
}

func ExtractIdempotencyKey(ctx context.Context) string {
	key, ok := ctx.Value(IdempotencyKey).(string)
	if !ok {
		panic("can not get idempotency key from context")
	}
	return key
}
