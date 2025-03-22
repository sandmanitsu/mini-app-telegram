package events

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (e *EventHandler) authMiddleware(update tgbotapi.Update) bool {
	if update.Message.Command() == startCmd {
		if e.userSrv.UserExist(update.Message.From.ID) {
			message := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! "+update.Message.From.FirstName)
			message.ReplyMarkup = e.keyboard()
			e.bot.Send(message)

			return false
		}

		e.createUser(update.Message)

		return false
	}

	if !e.userSrv.UserExist(update.Message.From.ID) {
		e.sendMessage(update.Message.Chat.ID, "Зарегистрируйтесь! команда /start")

		return false
	}

	return true
}
