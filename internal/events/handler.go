package events

import (
	"fmt"
	"mini-app-telegram/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCmd = "start"

	profileCmd = "Данные профиля"
)

type UserService interface {
	UserExist(userId int64) bool
	CreateUser(user domain.User)
	GetUser(userId int64) domain.User
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
	if update.Message.Command() == startCmd {
		e.userSrv.CreateUser(domain.User{
			UserId:    update.Message.From.ID,
			Username:  update.Message.From.UserName,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
		})

		message := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! "+update.Message.From.FirstName)
		message.ReplyMarkup = e.keyboard()
		e.bot.Send(message)

		return
	}

	if !e.authMiddleware(update) {
		return
	}

	e.doCommand(update.Message)
}

func (e *EventHandler) doCommand(msg *tgbotapi.Message) {
	fmt.Println(msg.Command())
	switch msg.Text {
	case profileCmd:
		user := e.userSrv.GetUser(msg.From.ID)

		e.sendMessage(msg.Chat.ID, fmt.Sprintf("Id: %d,\nUsername: %s,\nFirst Name: %s,\nLast Name: %v", user.UserId, user.Username, user.FirstName, user.LastName))
	default:
		e.sendMessage(msg.Chat.ID, "Неизвестная команда!")
	}
}

func (e *EventHandler) sendMessage(chatId int64, msg string) {
	message := tgbotapi.NewMessage(chatId, msg)
	e.bot.Send(message)
}

func (e *EventHandler) keyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(profileCmd),
		),
	)

	return keyboard
}
