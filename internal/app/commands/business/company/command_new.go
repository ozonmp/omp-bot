package company

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *CompanyCommander) New(inputMessage *tgbotapi.Message) {
	idx, err := c.companyService.Create(business.Company{})
	if err != nil {
		log.Printf("fail to create company: %v", err)
		return
	}

	outMessage := fmt.Sprintf("Created new company with idx %d", idx)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outMessage,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.Get: error sending reply message to chat - %v", err)
	}
}
