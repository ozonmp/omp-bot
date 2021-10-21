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

func (c *CompanyCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	split_args := strings.Split(args, " ")

	idx, err := strconv.Atoi(split_args[0])
	if err != nil {
		log.Println("CompanyCommander.Get: wrong idx arg", split_args[0])
		return
	}

	json_company := strings.Join(split_args[1:], " ")

	company := business.Company{}
	err = json.Unmarshal([]byte(json_company), &company)
	if err != nil {
		log.Printf("CompanyCommander.Edit: unable to unmarashal '%v': %v", json_company, err)
		return
	}

	err = c.companyService.Update(uint64(idx), company)
	if err != nil {
		log.Printf("CompanyCommander.Edit: fail to edit company with idx %d: %v", idx, err)
		return
	}
	outMessage := fmt.Sprintf("Updated company with idx %d.", idx)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outMessage,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.Edit: error sending reply message to chat - %v", err)
	}
}
