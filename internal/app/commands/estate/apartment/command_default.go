package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyApartmentCommander) Default(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	return
}
