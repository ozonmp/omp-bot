package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyApartmentCommander) Help(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	resp = tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products",
	)
	return
}
