package internship

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkInternshipCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please provide an id of internship!")
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkInternshipCommander.Delete: error sending reply message to chat - %v", err)
		}
		return
	}
	delResult, err := c.internshipService.Remove(uint64(idx))
	if err != nil {
		log.Printf("WorkInternshipCommander.Delete: error during remove internship from service - %v", err)
	}
	var msgText string
	if delResult {
		msgText = "Internship info is deleted!"
	} else {
		msgText = "Internship info with this id is not founded."
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.Delete: error sending reply message to chat - %v", err)
	}
}
