package main

import (
	"log"
	"log/slog"
	"mini-app-telegram/internal/app"
	"mini-app-telegram/internal/config"
	"mini-app-telegram/internal/logger"
)

func main() {
	log.Println("config initializing...")
	config := config.MustLoad()

	log.Println("logger initializing...")
	logger := logger.NewLogger(config.Env)
	logger.Info("logger started!", slog.String("env", config.Env))

	app.Run(config, logger)
}
