package test

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EducationTestCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	splitargs := strings.Fields(args)

	idx, err := strconv.Atoi(splitargs[0])
	if err != nil || len(splitargs) != 2 {
		log.Println("wrong args, pls write id and title", args)
		return
	}

	err = c.testService.Edit(idx, splitargs[1])
	if err != nil {
		log.Printf("fail to edit test with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"succesfully edited",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationTestCommander.Get: error sending reply message to chat - %v", err)
	}
}
