package company

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *CompanyCommander) New(inputMessage *tgbotapi.Message) {
	json_company := inputMessage.CommandArguments()

	company := business.Company{}
	err := json.Unmarshal([]byte(json_company), &company)
	if err != nil {
		log.Printf("CompanyCommander.New: unable to unmarashal '%v': %v", json_company, err)
		return
	}

	idx, err := c.companyService.Create(company)
	if err != nil {
		log.Printf("CompanyCommander.New: fail to create company: %v", err)
		return
	}

	outMessage := fmt.Sprintf("Created new company with idx %d", idx)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outMessage,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.New: error sending reply message to chat - %v", err)
	}
}
