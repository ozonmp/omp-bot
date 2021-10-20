package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyApartmentCommander) Default(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	resp = tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Sorry, I don't know command: "+inputMessage.Text+
			"\nHere is a list of available commands:\n"+helpText,
	)
	return
}
