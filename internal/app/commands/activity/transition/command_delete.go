package transition

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityTransitionCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	result, err := c.transitionService.Remove(uint64(idx))
	if err != nil {
		log.Printf("Fail to get transition with id %d: %v", idx, err)
		return
	}

	var txtResult string
	if result {
		txtResult = "Successful"
	} else {
		txtResult = "Failed"
	}
	textMsg := fmt.Sprintf("%s delete transition", txtResult)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		textMsg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityTransitionCommander.Delete: error sending reply message to chat - %v", err)
	}
}
