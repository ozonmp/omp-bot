package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ProductCommanderImpl) List(inputMessage *tgbotapi.Message) {
	c.sendPage(inputMessage.Chat.ID, 1)
}
