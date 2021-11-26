package product

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (commander *ProductCommander) Default(inputMessage *tgbotapi.Message) {
	commander.Send(inputMessage.Chat.ID, `"command not found.\n"
			"Use /help__recommendation__product to get commands list`,
	)
}
