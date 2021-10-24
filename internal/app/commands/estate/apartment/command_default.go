package apartment

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummyApartmentCommander) Default(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	respText := fmt.Sprintf("Sorry, I don't know command: %s\nHere is a list of available commands:\n%s", inputMessage.Text, helpText)
	resp = tgbotapi.NewMessage(inputMessage.Chat.ID, respText)
	return
}
