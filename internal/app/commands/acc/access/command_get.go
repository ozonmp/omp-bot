package access

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AccAccessCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	var outMsgText string
	access, err := c.accessService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get access with idx %d: %v", idx, err)
		outMsgText = "fail to get access"
	} else {
		outMsgText = c.accessService.String(*access)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outMsgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("AccAccessCommander.Get: error sending reply message to chat - %v", err)
	}
}
