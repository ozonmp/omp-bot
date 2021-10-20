package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpText = "/help__estate__apartment - help\n" +
	"/list__estate__apartment - list apartments (via pagination)\n" +
	"/get__estate__apartment <id> - get apartment with <id>\n" +
	"/delete__estate__apartment <id> - delete apartment with <id>\n"

func (c *DummyApartmentCommander) Help(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	resp = tgbotapi.NewMessage(inputMessage.Chat.ID, helpText)
	return
}
