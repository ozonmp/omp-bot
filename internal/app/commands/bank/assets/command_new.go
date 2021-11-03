package assets

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (c *AssetsCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	commands := strings.Split(args, " ")

	if len(commands) != 2 {
		c.Send(inputMessage.Chat.ID, "Неверное количество аргументов")
		return
	}

	userId, err := strconv.ParseUint(commands[0], 10, 64)
	if err != nil {
		c.Send(inputMessage.Chat.ID, "Проблемы со считыванием userId")
		return
	}
	price, err := strconv.ParseFloat(commands[1], 64)
	if err != nil || price < 0 {
		c.Send(inputMessage.Chat.ID, "Проблемы со считыванием цены, будет установлена цена = 0")
		price = 0
	}

	newID := c.assetsService.Create(userId, price)

	c.Send(inputMessage.Chat.ID, fmt.Sprintf("Актив %d успешно добавлен, ID=%d, price=%.2f", userId, newID, price))
}
