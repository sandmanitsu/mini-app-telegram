package events

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type MessageBuilder struct {
	message tgbotapi.MessageConfig
}

func NewMessageBuilder(chatId int64, text string) *MessageBuilder {
	return &MessageBuilder{
		message: tgbotapi.NewMessage(chatId, text),
	}
}

func (m MessageBuilder) setReplyKeyboard() MessageBuilder {
	m.message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(profileCmd),
		),
	)

	return m
}

func (m MessageBuilder) build() tgbotapi.MessageConfig {
	return m.message
}
