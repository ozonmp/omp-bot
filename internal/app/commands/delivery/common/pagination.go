package common

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

var pageSize = 3

func generatePaginationButtons(cursor int, c *DummyCommonCommander, msg *tgbotapi.MessageConfig) {
	var buttons []tgbotapi.InlineKeyboardButton

	previousOffset := cursor - pageSize
	hasPreviousPage := true
	if previousOffset < 0 {
		previousOffset = 0
		hasPreviousPage = false
	}

	if hasPreviousPage {
		buttons = append(
			buttons,
			paginationButton("Previous page", previousOffset),
		)
	}

	nextOffset := cursor + pageSize
	hasNextPage := true
	if _, err := c.commonService.List(uint64(nextOffset), 1); err != nil {
		nextOffset = cursor
		hasNextPage = false
	}

	if hasNextPage {
		buttons = append(
			buttons,
			paginationButton("Next page", nextOffset),
		)
	}

	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(buttons...),
		)
	}
}

func paginationButton(text string, offset int) tgbotapi.InlineKeyboardButton {
	callbackData, _ := json.Marshal(CallbackListData{
		Offset: offset,
	})

	callbackPath := path.CallbackPath{
		Domain:       "delivery",
		Subdomain:    "common",
		CallbackName: "list",
		CallbackData: string(callbackData),
	}

	return tgbotapi.NewInlineKeyboardButtonData(text, callbackPath.String())
}
