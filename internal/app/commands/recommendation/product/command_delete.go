package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (commander *ProductCommander) Delete(inputMessage *tgbotapi.Message) {
    args := inputMessage.CommandArguments()
	productId, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf(err.Error())
        commander.Send(inputMessage.Chat.ID,
			fmt.Sprintf("Not valid argument \"%v\"", args))
		return
	}
	if _, err := commander.service.Remove(productId); err != nil{
	    commander.Send(inputMessage.Chat.ID,
			fmt.Sprintf("Error: %s", err.Error()))
		log.Printf(err.Error())
	} else {
		commander.Send(inputMessage.Chat.ID, "Ok")
	}
}
