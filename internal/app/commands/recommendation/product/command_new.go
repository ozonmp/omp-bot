package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (commander *ProductCommander) New(inputMessage *tgbotapi.Message) {
    args := inputMessage.CommandArguments()
	productToCreate, err := commander.serializer.serialize(args)
	if err != nil{
		commander.Send(inputMessage.Chat.ID, fmt.Sprintf("Not valid argument \"%v\"", args))
		log.Printf(err.Error())
		return
	}
	_, err = commander.service.Create(productToCreate)
	if err != nil {
		commander.Send(inputMessage.Chat.ID, err.Error())
	} else {
		commander.Send(inputMessage.Chat.ID, "Ok")
	}
}
