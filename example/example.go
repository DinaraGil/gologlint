package main

import (
	"fmt"
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	slog.Info("hello 😝🫵")
	var a, b, c int
	var str string = "aaa"
	fmt.Println(a, b, c)
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	sugar.Info("Failed to fetch URL: " + str)
	slog.Error("Приветик")
	slog.Warn("hellll!!!!! ")
	slog.Warn("баобаб")
	slog.Info("my password is 123")
	slog.Info("my password is 123 %s", str)
	sugar.Infof("my password is 123 %s", str)
	fmt.Println("abc")
	sugar.Errorf("abc")
	Str1 := "Abc"
	slog.Info(Str1)
	slog.Info(str + "абв")
	slog.Info("")
	token := "aaa"
	slog.Info("token: " + token)
	log.Fatal("normal message")
}
