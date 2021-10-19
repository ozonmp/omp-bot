package ground

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroundCommander) Get(inputMessage *tgbotapi.Message) {
	var msgText string

	defer func() {
		c.Send(
			inputMessage.Chat.ID,
			msgText,
		)
	}()

	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("Internal error %v", err)
		msgText = fmt.Sprintf("Wrong args `%s`. Id > 0", args)
		return
	}

	product, err := c.service.Describe(idx)
	if err != nil {
		log.Printf("Internal error %v", err)
		msgText = fmt.Sprintf("Fail to get ground with id %d", idx)
	} else {
		msgText = product.String()
	}
}
