package product

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ProductCommanderImpl) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	isDel, err := c.service.Remove(uint64(idx))
	if err != nil || !isDel {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"The product has been deleted",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductCommanderImpl.Delete: error sending reply message to chat - %v", err)
	}
}
