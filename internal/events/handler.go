package events

import (
	"mini-app-telegram/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCmd = "start"

	// keyboard commands
	profileCmd = "Данные профиля"
)

type UserService interface {
	CreateUser(user domain.User) error
	GetUser(userId int64) (domain.User, error)
	UserExist(userId int64) bool
}

type EventHandler struct {
	bot     *tgbotapi.BotAPI
	userSrv UserService
}

func NewEventHandler(bot *tgbotapi.BotAPI, userSrv UserService) *EventHandler {
	return &EventHandler{
		bot:     bot,
		userSrv: userSrv,
	}
}

func (e *EventHandler) Handle(update tgbotapi.Update) {
	if !e.authMiddleware(update) {
		return
	}

	e.doCommand(update.Message)
}

func (e *EventHandler) doCommand(msg *tgbotapi.Message) {
	switch msg.Text {
	case profileCmd:
		e.getProfile(msg)
	default:
		e.bot.Send(
			NewMessageBuilder(msg.Chat.ID, "Неизвестная команда!").build(),
		)
	}
}
