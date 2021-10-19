package project

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/work"
	"log"
)

func (c *ProjectCommander) New(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	var project work.Project
	err := json.Unmarshal([]byte(args), &project)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	idx, err := c.projectService.Create(project)
	if err != nil {
		log.Println("failed to create product")
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("created project %v", idx),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Get: error sending reply message to chat - %v", err)
		return
	}

}
