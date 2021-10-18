package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (command *ProductCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__recomendation__product — print list of commands\n"+
		     "/get__recomendation__product — get product by id\n"+
			 "/list__recomendation__product - list of all products\n" +
			 "/delete__recomendation__product - delete product by id\n" +
			 "/new__recomendation__product - create new product\n" +
			 "/edit__recomendation__product - edit product\n",
	)
	command.bot.Send(msg)
}

