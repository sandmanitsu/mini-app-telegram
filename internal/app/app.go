package app

import (
	"log/slog"
	"mini-app-telegram/internal/config"
	"mini-app-telegram/internal/events"
	repository "mini-app-telegram/internal/repository/user"
	"mini-app-telegram/internal/service/user"
	"mini-app-telegram/internal/storage/postgresql"
	"os"

	sl "mini-app-telegram/internal/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	envDev = "dev"
)

func Run(config *config.Config, logger *slog.Logger) {
	// storage
	storage, err := postgresql.NewPostgreSQL(config.DB)
	if err != nil {
		logger.Error("failed to init postgresql storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage

	// repository init
	userRepo := repository.NewUserRepository(storage.DB, logger)

	// services init
	userSrv := user.NewUserService(userRepo)

	// start bot
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		logger.Error("failed to create bot instanse", sl.Err(err))
		os.Exit(1)
	}

	if config.Env == envDev {
		bot.Debug = true
	}

	logger.Info("bot authorized!", slog.String("env", config.Env))

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	handler := events.NewEventHandler(bot, userSrv)

	for update := range updates {
		handler.Handle(update)
	}
}
