package order

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *OrderCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil || id < 0 {
		log.Printf("OrderCommander.Get "+
			"error parsing id: %v", err)

		c.Reply(
			inputMessage.Chat.ID,
			"Failed to parse order id! Correct syntax for 'get' command is:\n"+
				`/get__buy__order <id> (id >= 0)`)
		return
	}

	order, err := c.orderService.Describe(uint64(id))
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", id, err)

		c.Reply(
			inputMessage.Chat.ID,
			fmt.Sprintf(`Failed to get order: %v`, err))
		return
	}

	c.Reply(
		inputMessage.Chat.ID,
		order.String())
}
