package production

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	listLimit          = 3
	listPaginationText = "More"
	listHeader         = "Id Description Type"
)

const (
	listCallbackDomain    = "recommendation"
	listCallbackSubdomain = "production"
	listCallbackName      = "list"
)

func (c *RecommendationProductionCommander) List(inputMessage *tgbotapi.Message) {
	c.list(inputMessage.Chat.ID, CallbackListData{})
}

func (c *RecommendationProductionCommander) list(chatId int64, data CallbackListData) {
	text, ln, err := func() (string, int, error) {
		productions, err := c.productionService.List(data.Offset, listLimit)
		if err != nil {
			return textWrong, 0, err
		}

		if len(productions) == 0 && data.MessageId != 0 {
			return "", 0, nil
		}

		var builder strings.Builder
		if data.MessageId == 0 {
			if _, err := fmt.Fprintf(&builder, "%v\n", listHeader); err != nil {
				return textWrong, 0, err
			}
		}

		for _, v := range productions {
			if _, err := fmt.Fprintf(&builder, "%v\n", v); err != nil {
				return textWrong, 0, err
			}
		}

		return builder.String(), len(productions), nil
	}()
	if err != nil {
		log.Println(err)
	}

	if data.MessageId != 0 {
		c.deleteListMarkup(chatId, data)
	}

	if text == "" {
		return
	}

	msg := tgbotapi.NewMessage(chatId, text)

	retMsg, err := c.sendMessage(msg)
	if err != nil {
		return
	}

	if ln == listLimit {
		data.MessageId = retMsg.MessageID
		c.listMarkup(chatId, data)
	}
}

func (c *RecommendationProductionCommander) listMarkup(chatId int64, data CallbackListData) {
	j, _ := json.Marshal(data)

	callbackPath := path.CallbackPath{
		Domain:       listCallbackDomain,
		Subdomain:    listCallbackSubdomain,
		CallbackName: listCallbackName,
		CallbackData: string(j),
	}

	markupMsg := tgbotapi.NewEditMessageReplyMarkup(chatId, data.MessageId,
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(listPaginationText, callbackPath.String()),
			),
		),
	)

	c.sendMessage(markupMsg)
}

func (c *RecommendationProductionCommander) deleteListMarkup(chatId int64, data CallbackListData) {
	markupMsg := tgbotapi.NewEditMessageReplyMarkup(chatId, data.MessageId,
		tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{}),
	)

	c.sendMessage(markupMsg)
}
