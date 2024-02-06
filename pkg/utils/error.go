package utils

import (
	"log/slog"
	"os"
)

func Fatal(msg string, err error) {
	slog.Error(msg, slog.Any("err", err.Error()))
	os.Exit(1)
}

func Panic(err error) {
	slog.Error("err", slog.Any("err", err.Error()))
	panic(err)
}
