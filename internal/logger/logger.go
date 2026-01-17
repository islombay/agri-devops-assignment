package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	Env      string
	LogLevel string
}

func New(cfg Config) *slog.Logger {
	var level slog.Level

	switch cfg.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	handlerOpts := &slog.HandlerOptions{
		Level: level,
	}

	var handler slog.Handler

	if cfg.Env == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, handlerOpts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, handlerOpts)
	}

	return slog.New(handler)
}
