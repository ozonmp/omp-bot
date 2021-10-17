package announcement

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *StreamingAnnouncementCommander) New(inputMessage *tgbotapi.Message) {
	// TODO
	c.Default(inputMessage)
}
