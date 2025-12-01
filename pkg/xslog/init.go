package xslog

import (
	"log/slog"
	"os"
)

func init() {
	lvl := new(slog.Level)
	_ = lvl.UnmarshalText([]byte(os.Getenv("LOG_LEVEL")))

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	})

	slog.SetDefault(slog.New(handler))
}
