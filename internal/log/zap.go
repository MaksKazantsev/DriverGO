package log

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	*zap.SugaredLogger
}

func (z zapLogger) Error(msg string, data *Data) {
	if data == nil {
		z.Errorw(msg)
	} else {
		z.Errorw(msg, data.Key, data.Val)
	}
}

func (z zapLogger) Info(msg string, data *Data) {
	if data == nil {
		z.Infow(msg)
	} else {
		z.Infow(msg, data.Key, data.Val)
	}
}

func (z zapLogger) Trace(key string, msg string) {
	z.Infow(msg, "key", key)
}

var _ Logger = &zapLogger{}

func newLocalLogger() *zapLogger {
	l, err := zap.NewProduction()
	if err != nil {
		panic("failed to init logger: " + err.Error())
	}
	return &zapLogger{l.Sugar()}
}
