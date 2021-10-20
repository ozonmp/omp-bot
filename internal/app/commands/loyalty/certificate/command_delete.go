package certificate

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *LoyaltyCertificateCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	found, err := c.certificateService.Remove(uint64(id))
	if err != nil {
		log.Printf("failed to delete certificate with id %d: %v", id, err)
		return
	}

	outputMsg := "Certificate with ID " + strconv.Itoa(id)
	if found {
		outputMsg += " was removed"
	} else {
		outputMsg += " not found"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.Get: error sending reply message to chat - %v", err)
	}
}
