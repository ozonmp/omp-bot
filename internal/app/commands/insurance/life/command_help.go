package life

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (telegramLifeCommander *TelegramLifeCommander) Help(inputMessage *tgbotapi.Message) {

	responseMessage :=
		"/help - help\n" +
			"list - list products"

	telegramLifeCommander.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, responseMessage))

}
