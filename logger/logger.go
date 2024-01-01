package logger

import (
	"log/slog"
	"os"
)

func LogInfo() *slog.Logger {
	var log *slog.Logger
	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}

func LogErr() *slog.Logger {
	var log *slog.Logger
	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	return log
}
