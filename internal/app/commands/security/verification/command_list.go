package verification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const MaxUint = ^uint64(0)

func (c *SecurityVerificationCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	//list should return all values
	products, err := c.verificationService.List(0, MaxUint)
	if err != nil {
		log.Printf("fail to get list of products: %v", err)
		c.sendErrorMsg("List", tgbotapi.NewMessage(inputMessage.Chat.ID, internalError))
		return
	}
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}
