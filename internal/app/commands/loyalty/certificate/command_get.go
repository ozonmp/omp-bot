package certificate

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *LoyaltyCertificateCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	if id < 0 {
		log.Println("id must be positive", args)
		return
	}

	certificate, err := c.certificateService.Describe(uint64(id))
	if err != nil {
		log.Printf("fail to get certificate with id %d: %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("%s\n", certificate.String()),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.Get: error sending reply message to chat - %v", err)
	}
}
