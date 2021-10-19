package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (commander *ProductCommander) Get(inputMessage *tgbotapi.Message) {
    args := inputMessage.CommandArguments()
	productId, err := strconv.ParseUint(args, 10, 64)
	if err != nil{
		commander.Send(inputMessage.Chat.ID,
		fmt.Sprintf("Not valid argument \"%v\"", args))
		log.Printf("ProductCommander.Get: %s", err.Error())
		return
	}
	product, err := commander.service.Describe(productId)
	if err != nil{
		commander.Send(inputMessage.Chat.ID,
			fmt.Sprintf("Error: %s", err.Error()))
		log.Printf("ProductCommander.Get: %s", err.Error())
		return
	}

	text, err := commander.serializer.deserialize(product)
	if err != nil{
		text = "Error in data serializing"
		log.Printf("ProductCommander.Get: %s", err.Error())
	}
	commander.Send(inputMessage.Chat.ID, text)
}
