package assets

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AssetsCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Неверный аргумент %s: %v", args, err))
		return
	}

	err = c.assetsService.Remove(idx)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Не смог удалить актив %d: %v", idx, err))
		return
	}

	c.Send(inputMessage.Chat.ID, fmt.Sprintf("Успешно удалён элемент ID=%d", idx))
}
