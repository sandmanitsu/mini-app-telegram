package events

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (e *EventHandler) authMiddleware(update tgbotapi.Update) bool {
	if update.Message.Command() == startCmd {
		if e.userSrv.UserExist(update.Message.From.ID) {
			e.bot.Send(
				NewMessageBuilder(update.Message.Chat.ID, "Привет! "+update.Message.From.FirstName).
					setReplyKeyboard().
					build(),
			)

			return false
		}

		e.createUser(update.Message)

		return false
	}

	if !e.userSrv.UserExist(update.Message.From.ID) {
		e.bot.Send(
			NewMessageBuilder(update.Message.Chat.ID, "Зарегистрируйтесь! команда /start").build(),
		)

		return false
	}

	return true
}
