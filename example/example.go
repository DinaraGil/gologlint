package main

import (
	"fmt"
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	slog.Info("hello")
	var a, b, c int
	fmt.Println(a, b, c)
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	sugar.Info("Failed to fetch URL: %d", a)
	slog.Error("hellll")
	slog.Warn("hellll")
	log.Fatal("aaa")
}
