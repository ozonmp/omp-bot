package customer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/rating/customer/paginator"
)

func (c *CustomerCommander) List(inputMessage *tgbotapi.Message) error {
	msg, err := c.paginator.GetMessage(inputMessage.Chat.ID, paginator.CallbackListData{})
	if err != nil {
		return err
	}
	_, err = c.bot.Send(msg)
	return err
}
