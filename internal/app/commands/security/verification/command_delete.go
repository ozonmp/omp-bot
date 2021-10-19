package verification

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SecurityVerificationCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	isDeleted, err := c.verificationService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to delete product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"product with idx "+args+" was deleted",
	)

	if !isDeleted {
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"product with idx "+args+" was not deleted",
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("SecurityVerificationCommander.Delete: error sending reply message to chat - %v", err)
	}
}
