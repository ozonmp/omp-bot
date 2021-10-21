package company

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *CompanyCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	company := business.Company{}

	if args[:1] == "{" {
		err := json.Unmarshal([]byte(args), &company)
		if err != nil {
			log.Printf("CompanyCommander.New: fail to create company, invalid json stucture")
			return
		}
	} else {
		err := readPlain(args, &company)
		if err != nil {
			log.Printf("CompanyCommander.New: fail to create company, %v", err)
			return
		}
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

func readPlain(args string, company *business.Company) error {
	arg_company := strings.Split(args, "|")

	if len(arg_company) != 3 {
		return fmt.Errorf("invalid stucture zipcode|name|address")
	}

	company_zipcode, err := strconv.Atoi(arg_company[0])
	if err != nil {
		return fmt.Errorf("CompanyCommander.New: fail to create company, invalid stucture zipcode|name|address : %v", err)
	}

	company.ZipCode = int64(company_zipcode)
	company.Name = arg_company[1]
	company.Address = arg_company[2]

	return nil
}
