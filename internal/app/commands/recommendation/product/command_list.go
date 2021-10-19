package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ProductCommander) List(inputMessage *tgbotapi.Message) {
	commander.ShowButtons(0, 2, inputMessage.Chat.ID)
}
