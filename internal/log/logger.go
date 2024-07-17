package log

type Logger interface {
	Error(msg string, data *Data)
	Info(msg string, data *Data)
	Trace(key string, msg string)
}

type Data struct {
	Key string
	Val any
}

func WithData(key string, val any) *Data {
	return &Data{Key: key, Val: val}
}

func MustInit(t string) Logger {
	switch t {
	case "local":
		return newLocalLogger()
	default:
		panic("unknown logger type")
	}
}
