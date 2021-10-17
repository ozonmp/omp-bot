package service

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/service/service"
)

func (c *ServiceServiceCommander) Delete(inputMessage *tgbotapi.Message) (tgbotapi.Message, error) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		return tgbotapi.Message{}, service.NewBadRequestError(fmt.Sprintf("Wrong args %v", args))
	}

	_, err = c.serviceService.Remove(idx)
	if err != nil {
		return tgbotapi.Message{}, service.NewBadRequestError(fmt.Sprintf("Failed to remove productwith idx %d: %v", idx, err))
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Product with index %d was removed", idx),
	)

	return c.bot.Send(msg)
}
