package subdomain

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *DummySubdomainCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	pagination := NewPaginationList(c.subdomainService.List, parsedData.Cursor, parsedData.Limit)

	page := pagination.Page()
	buttons := pagination.Buttons()
	inlineKeyboardButtons := make([]tgbotapi.InlineKeyboardButton, 0, 2)
	for _, v := range buttons {
		inlineKeyboardButtons = append(
			inlineKeyboardButtons,
			tgbotapi.NewInlineKeyboardButtonData(v.Text, v.Data),
		)
	}
	keyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(inlineKeyboardButtons...),
	)

	editConf := tgbotapi.EditMessageTextConfig{
		Text: page,
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:    callback.Message.Chat.ID,
			MessageID: callback.Message.MessageID,
		},
	}
	if len(buttons) > 0 {
		editConf.BaseEdit.ReplyMarkup = &keyboardMarkup
	}

	c.bot.Send(editConf)
}
