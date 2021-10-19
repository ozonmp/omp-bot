package workplace

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessWorkplaceCommander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	c.bot.Send(msg)
}
