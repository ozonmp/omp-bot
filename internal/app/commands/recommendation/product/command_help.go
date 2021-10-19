package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (commander *ProductCommander) Help(inputMessage *tgbotapi.Message) {
    jsonFormat := 		"Data in JSON format:\n" +
		"{\n" +
		"  \"Id\": number\n" +
		"  \"Title\":\"string\"\n" +
		"  \"Description\":\"string\"\n" +
		"  \"Rating\" : number\n" +
		"}\n"
	commander.Send(inputMessage.Chat.ID,
		"/help__recommendation__product  id — print list of commands\n"+
		     "/get__recommendation__product id — get product by id\n"+
			 "/list__recommendation__product - list of all products\n" +
			 "/delete__recommendation__product id- delete product by id\n" +
			 "/new__recommendation__product data - create new product\n" + jsonFormat +
			 "/edit__recommendation__product id data - edit product by id\n" + jsonFormat,
	)
}

