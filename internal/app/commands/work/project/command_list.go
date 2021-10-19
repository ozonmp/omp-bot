package project

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *ProjectCommander) List(inputMessage *tgbotapi.Message) {

	projectList, err := c.projectService.List(0, 0)
	if err != nil {
		log.Println("fail to get product list")
		return
	}

	text := ""
	for _, v := range projectList {
		text += v.String() + "\n"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		text,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Get: error sending reply message to chat - %v", err)
		return
	}
}
