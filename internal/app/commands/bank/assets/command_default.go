package assets

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AssetsCommander) Default(inputMessage *tgbotapi.Message) {
	c.Send(inputMessage.Chat.ID, "Неизвестная команда")
	c.Help(inputMessage)
}
