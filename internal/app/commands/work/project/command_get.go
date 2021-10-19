package project

import (


	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *ProjectCommander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	project, err := c.projectService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		project.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Get: error sending reply message to chat - %v", err)
		return
	}
}
