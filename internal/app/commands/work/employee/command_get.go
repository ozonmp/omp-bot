package employee

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Get(inputMessage *tgbotapi.Message) {
	msgText := ""
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		msgText = "Error: wrong args"
	} else {
		msgText = c.subdomainService.Get(idx)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}
}
