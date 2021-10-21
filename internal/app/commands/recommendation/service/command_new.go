package service

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/service/recommendation/service"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RecommendationServiceCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	splitArgs := strings.Split(args, " ")
	var messageText string
	if len(splitArgs) == 2 {
		service := service.NewService(0, splitArgs[0], splitArgs[1])
		idx, err := c.serviceService.Create(*service)
		if err == nil {
			messageText = fmt.Sprintf("Created new service with Id = %d", idx)
		} else {
			messageText = fmt.Sprintf("Error: %s", err)
		}
	} else {
		log.Println("wrong args", args)
		messageText = "Please input two string(title and description) with space separator"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		messageText,
	)

	c.bot.Send(msg)
}
