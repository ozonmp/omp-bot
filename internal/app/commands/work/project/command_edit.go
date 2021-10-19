package project

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/work"
	"log"
)

func (c *ProjectCommander) Edit(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	var project work.Project
	err := json.Unmarshal([]byte(args), &project)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	err = c.projectService.Update(project.ID, project)
	if err != nil {
		log.Printf("fail to edit product with idx %d: %v", project.ID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"updated",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Get: error sending reply message to chat - %v", err)
		return
	}
}
