package events

import (
	"fmt"
	"mini-app-telegram/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (e *EventHandler) createUser(msg *tgbotapi.Message) {
	e.userSrv.CreateUser(domain.User{
		UserId:    msg.From.ID,
		Username:  msg.From.UserName,
		FirstName: msg.From.FirstName,
		LastName:  msg.From.LastName,
	})

	message := tgbotapi.NewMessage(msg.Chat.ID, "Привет! "+msg.From.FirstName)
	message.ReplyMarkup = e.keyboard()
	e.bot.Send(message)
}

func (e *EventHandler) getProfile(msg *tgbotapi.Message) {
	user := e.userSrv.GetUser(msg.From.ID)

	e.sendMessage(msg.Chat.ID, fmt.Sprintf("Id: %d,\nUsername: %s,\nFirst Name: %s,\nLast Name: %v", user.UserId, user.Username, user.FirstName, user.LastName))
}
