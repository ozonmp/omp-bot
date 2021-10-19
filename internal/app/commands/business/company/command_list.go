package company

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CompanyCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the companies: \n\n"

	companies := c.companyService.List()
	for _, c := range companies {
		outputMsgText += c.Name
		outputMsgText += " [zipcode" + strconv.Itoa(int(c.ZipCode)) + "]"
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.List: error sending reply message to chat - %v", err)
	}
}
