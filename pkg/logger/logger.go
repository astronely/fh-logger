package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local.env"
	envProd  = "prod.env"
	envDev   = "dev.env"
)

var globalLogger *slog.Logger

func Init(env string) {
	var opts *slog.HandlerOptions

	switch env {
	case envLocal:
		opts = &slog.HandlerOptions{Level: slog.LevelDebug}
	case envDev:
		opts = &slog.HandlerOptions{Level: slog.LevelDebug}
	case envProd:
		opts = &slog.HandlerOptions{Level: slog.LevelInfo}
	}

	prettyOpts := PrettyHandlerOptions{
		SlogOpts: opts,
	}
	handler := prettyOpts.NewPrettyHandler(os.Stdout)

	globalLogger = slog.New(handler)
}

func Info(msg string, args ...any) {
	globalLogger.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}

func Warn(msg string, args ...any) {
	globalLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	globalLogger.Error(msg, args...)
}

func With(args ...any) {
	globalLogger.With(args)
}
