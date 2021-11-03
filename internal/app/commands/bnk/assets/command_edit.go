package assets

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AssetsCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	commands := strings.Split(args, " ")

	for _, command := range commands {
		if command == "" {
			c.Send(inputMessage.Chat.ID, "Один из элементов пуст")
			return
		}
	}

	ID, err := strconv.ParseUint(commands[0], 10, 64)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Неверный аргумент %s: %v", ID, err))
		return
	}
	price, err := strconv.ParseFloat(commands[1], 64)
	if err != nil || price < 0 {
		c.Send(inputMessage.Chat.ID, "Проблемы со считыванием цены, будет установлена цена = 0")
		price = 0
	}

	err = c.assetsService.Update(ID, price)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Ошибки обновления актива: %v", err))
		return
	}

	c.Send(inputMessage.Chat.ID, fmt.Sprintf("Актив %d успешно изменен на %.2f", ID, price))
}
