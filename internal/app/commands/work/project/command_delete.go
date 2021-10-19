package project

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *ProjectCommander) Delete(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	ok, err := c.projectService.Remove(uint64(idx))
	if err != nil || !ok {
		log.Printf("fail to remove product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("project %v is deleted", idx),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("Get: error sending reply message to chat - %v", err)
		return
	}

}
