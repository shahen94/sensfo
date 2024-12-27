package logger

import (
	"log/slog"
	"os"
)

type SlogLogger struct {
	instance *slog.Logger
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.instance.Info(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.instance.Error(msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	l.instance.Warn(msg, args...)
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.instance.Debug(msg, args...)
}

func NewSlogLogger() Logger {
	slog := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return &SlogLogger{
		instance: slog,
	}
}
