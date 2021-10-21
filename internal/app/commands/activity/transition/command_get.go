package transition

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityTransitionCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	result, err := c.transitionService.Describe(uint64(idx))
	if err != nil {
		log.Printf("Fail to get transition with idx %d: %v", idx, err)
		return
	}

	textMsg := fmt.Sprintf("%v", result)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		textMsg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityTransitionCommander.Get: error sending reply message to chat - %v", err)
	}
}
