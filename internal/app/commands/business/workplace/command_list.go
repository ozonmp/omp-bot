package workplace

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


func (c *BusinessWorkplaceCommander) List(inputMessage *tgbotapi.Message) {
	c.processList(0, inputMessage.Chat.ID)
}

