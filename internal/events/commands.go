package events

import (
	"fmt"
	"mini-app-telegram/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (e *EventHandler) createUser(msg *tgbotapi.Message) {
	err := e.userSrv.CreateUser(domain.User{
		UserId:    msg.From.ID,
		ChatId:    msg.Chat.ID,
		Username:  msg.From.UserName,
		FirstName: msg.From.FirstName,
		LastName:  msg.From.LastName,
	})
	if err != nil {
		e.bot.Send(
			NewMessageBuilder(msg.Chat.ID, errMsgCreateUser).build(),
		)

		return
	}

	message := NewMessageBuilder(msg.Chat.ID, "Привет! "+msg.From.FirstName).
		setReplyKeyboard().
		build()

	e.bot.Send(message)
}

func (e *EventHandler) getProfile(msg *tgbotapi.Message) {
	user, err := e.userSrv.GetUser(msg.From.ID)
	if err != nil {
		e.bot.Send(
			NewMessageBuilder(msg.Chat.ID, errMsgGetProfile).build(),
		)
	}

	text := fmt.Sprintf("Id: %d,\nUsername: %s,\nFirst Name: %s,\nLast Name: %v", user.UserId, user.Username, user.FirstName, user.LastName)

	e.bot.Send(
		NewMessageBuilder(msg.Chat.ID, text).build(),
	)
}
