package service

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/service/service"
)

func (c *ServiceServiceCommander) Get(inputMessage *tgbotapi.Message) (tgbotapi.Message, error) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		return tgbotapi.Message{}, service.NewBadRequestError(fmt.Sprintf("Wrong args %v", args))
	}

	entry, err := c.serviceService.Get(idx)
	if err != nil {
		return tgbotapi.Message{}, service.NewBadRequestError(fmt.Sprintf("Failed to get product with idx %d: %v", idx, err))
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		entry.String(),
	)

	return c.bot.Send(msg)
}
