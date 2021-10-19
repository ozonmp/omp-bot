package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (commander *ProductCommander) Edit(inputMessage *tgbotapi.Message) {
	args := strings.SplitAfterN(inputMessage.CommandArguments(), " ", 2)
	if len(args) < 2 {
		msg := fmt.Sprintf("Not valid argument \"%v\"", args)
		commander.Send(inputMessage.Chat.ID, msg)
		log.Printf("ProductCommander.Edit: %s", msg)
		return
	}
	productId, err := strconv.ParseUint(strings.Trim(args[0], " "), 10, 64)
	if err != nil {
		commander.Send(inputMessage.Chat.ID,
			fmt.Sprintf("Not valid argument \"%v\"", args))
		log.Printf("ProductCommander.Edit: %s", err.Error())
		return
	}
	productToEdit, err := commander.serializer.serialize(args[1])
	if err != nil {
		commander.Send(inputMessage.Chat.ID, fmt.Sprintf("Not valid argument \"%v\"", args))
		log.Printf(err.Error())
		return
	}
	err = commander.service.Update(productId, productToEdit)
	if err != nil {
		commander.Send(inputMessage.Chat.ID, err.Error())
		return
	}
	commander.Send(inputMessage.Chat.ID, "Ok")

}
