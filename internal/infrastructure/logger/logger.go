package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func Init(modeDebug bool) {
	var level slog.Level
	if modeDebug {
		level = slog.LevelDebug
	} else {
		level = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	Log = slog.New(handler)
}

func InitDebug() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	Log = slog.New(handler)
}
