package utils

type ContextKey string

const (
	LoggerKey      ContextKey = "loggerKey"
	IdempotencyKey ContextKey = "idemKey"
)
