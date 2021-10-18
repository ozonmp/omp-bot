package internship

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkInternshipCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please provide an id of internship!")
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkInternshipCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}
	internship, err := c.internshipService.Describe(uint64(idx))
	var msgText string
	if err != nil {
		log.Printf("fail to get internship with idx %d: %v", idx, err)
		msgText = "Can't find internship with this id :("
	} else {
		msgText = c.internshipService.FullString(*internship)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.Get: error sending reply message to chat - %v", err)
	}
}
