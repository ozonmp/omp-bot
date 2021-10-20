package warehouse

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EstateWarehouseCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__estate__warehouse - help\n"+
			"/list__estate__warehouse - list warehouses\n"+
			"/delete__estate__warehouse - delete warehouses\n"+
			"/new__estate__warehouse - new list warehouse\n"+
			"/get__estate__warehouse - get list warehouse\n"+
			"/edit__estate__warehouse - edit list warehouse\n",
	)

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/help__estate__warehouse"),
			tgbotapi.NewKeyboardButton("/list__estate__warehouse"),
			tgbotapi.NewKeyboardButton("/new__estate__warehouse"),
			tgbotapi.NewKeyboardButton("/edit__estate__warehouse"),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EstateWarehouseCommander.Help: error sending reply message to chat - %v", err)
	}
}
