package assets

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AssetsCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Неверный аргумент %s: %v", args, err))
		return
	}

	asset, err := c.assetsService.Get(idx)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Не найден элемент с ID=%d: %v", idx, err))
		return
	}

	c.Send(inputMessage.Chat.ID, asset.String())
}
