package company

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *CompanyCommander) New(inputMessage *tgbotapi.Message) {
	arg_company := strings.Split(inputMessage.CommandArguments(), "|")

	if len(arg_company) != 3 {
		log.Printf("CompanyCommander.New: fail to create company, invalid stucture zipcode|name|address")
		return
	}

	company_zipcode, err := strconv.Atoi(arg_company[0])
	if err != nil {
		log.Printf("CompanyCommander.New: fail to create company, invalid stucture zipcode|name|address : %v", err)
		return
	}

	company := business.Company{Name: arg_company[1], Address: arg_company[2], ZipCode: int64(company_zipcode)}

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
