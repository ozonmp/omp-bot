package solution

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service_consts"
	"strings"
)

func (c *SolutionCommander) List(inputMsg *tgbotapi.Message){
	cb := CallbackListData{ 1, 1}
	data, _ := json.Marshal(cb)

	callbackPath := path.CallbackPath{
		Domain:       service_consts.Education,
		Subdomain:    service_consts.Solution,
		CallbackName: "list",
		CallbackData: string(data),
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, strings.Join(c.SolutionService.List(cb.Start-1, cb.Offset), "\n"))
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next", callbackPath.String(),
			)))
	c.bot.Send(msg)
}

