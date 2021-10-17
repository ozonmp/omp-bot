package click

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *ActivityClickCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("ActivityClickCommander.Delete: error parsing product index: %v\n", err)

		c.SendMessageToChat(
			tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Error parsing product index: %v", err)),
			"ActivityClickCommander.Delete",
		)

		return
	}

	ok, err := c.service.Remove(idx)
	if err != nil {
		log.Printf("ActivityClickCommander.Delete: error deleting item by index %d: %v\n", idx, err)

		c.SendMessageToChat(
			tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Error: %s", err.Error())),
			"ActivityClickCommander.Delete",
		)

		return
	}

	if !ok {
		c.SendMessageToChat(
			tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Unable to remove item %d. Try again", idx)),
			"ActivityClickCommander.Delete",
		)

		return
	}

	c.SendMessageToChat(
		tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Item %d succesfully removed", idx)),
		"ActivityClickCommander.Delete",
	)
}
