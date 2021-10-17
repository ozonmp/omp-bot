package customer

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/rating/customer/paginator"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CustomerCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) error {
	parsedData := paginator.CallbackListData{}
	if err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData); err != nil {
		return err
	}

	msg, err := c.paginator.GetMessage(callback.Message.Chat.ID, parsedData)
	if err != nil {
		return err
	}
	_, err = c.bot.Send(msg)
	return err
}
