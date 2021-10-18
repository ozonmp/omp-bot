package internship

import (
	"github.com/ozonmp/omp-bot/internal/service/work/internship"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkInternshipCommander) New(inputMessage *tgbotapi.Message) {
	internshipID, err := c.internshipService.Create(internship.Internship{})
	var msgText string
	if err != nil {
		log.Printf("fail to create new internship %v", err)
		msgText = "Something wrong..."
	} else {
		msgText = "New internship ID: " + strconv.Itoa(int(internshipID)) + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.New: error sending reply message to chat - %v", err)
	}
}
