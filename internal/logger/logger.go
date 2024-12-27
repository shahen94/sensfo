package logger

type Logger interface {
	Info(string, ...any)
	Error(string, ...any)
	Warn(string, ...any)
	Debug(string, ...any)
}
