package workplace

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const pageSize uint64 = 3

func (c *BusinessWorkplaceCommander) List(inputMessage *tgbotapi.Message) {
	c.processList(0, pageSize, inputMessage.Chat.ID)
}

