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
	var str string = "aaa"
	fmt.Println(a, b, c)
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	sugar.Info("Failed to fetch URL: " + str)
	slog.Error("Hellll")
	slog.Warn("hellll")
	log.Fatal("aaa")
}
