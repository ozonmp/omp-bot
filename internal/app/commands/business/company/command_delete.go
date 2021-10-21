package company

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CompanyCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("CompanyCommander.Delete: wrong args", args)
		return
	}

	isDeleted, err := c.companyService.Remove(uint64(idx))
	if err != nil {
		log.Printf("CompanyCommander.Delete: unable to delete %d element - %v", idx, err)
	}

	if isDeleted {
		text := fmt.Sprintf("Company with index %d was deleted.\n", idx)

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			text,
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("CompanyCommander.Delete: error sending reply message to chat - %v", err)
		}
	}
}
