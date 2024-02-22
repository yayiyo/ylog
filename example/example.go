package main

import (
	"log"
	"os"

	"github.com/yayiyo/ylog"
)

func main() {
	ylog.Info("std log")
	ylog.SetOptions(ylog.WithLevel(ylog.LevelDebug))
	ylog.Debug("change std log to debug level")
	ylog.SetOptions(ylog.WithFormatter(&ylog.JsonFormatter{IgnoreBasicFields: false}))
	ylog.Debug("log in json format")
	ylog.Info("another log in json format")

	// 输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("create file test.log failed: %v", err)
	}

	defer fd.Close()
	l := ylog.New(ylog.WithLevel(ylog.LevelInfo),
		ylog.WithOutput(fd),
		ylog.WithFormatter(&ylog.JsonFormatter{IgnoreBasicFields: false}),
	)
	l.Info("custom log in json format")
}
