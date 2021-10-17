package production

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RecommendationProductionCommander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, textWrong)

	c.sendMessage(msg)
}
