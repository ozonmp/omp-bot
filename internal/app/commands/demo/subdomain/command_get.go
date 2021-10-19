package subdomain

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Get(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		err = fmt.Errorf("wrong args: %v", args)
		return
	}

	product, err := c.subdomainService.Get(idx)
	if err != nil {
		err = fmt.Errorf("failed to get product with idx %d: %v", idx, err)
		return
	}

	resp = tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	return
}
