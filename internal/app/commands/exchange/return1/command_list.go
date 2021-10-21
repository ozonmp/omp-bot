package return1

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/exchange/return1"
)

const PageSize = 2

type PaginationMarkup struct {
	Cursor uint64
	Limit  uint64
}

func (c *Return1CommanderImpl) List(inputMsg *tgbotapi.Message) {
	c.ListPaginated(inputMsg, 0, PageSize)
}

func (c *Return1CommanderImpl) ListPaginated(inputMsg *tgbotapi.Message, cursor uint64, limit uint64) {
	reply := func(text string, markup *tgbotapi.InlineKeyboardMarkup, other ...interface{}) {
		for _, v := range other {
			log.Println("Return1CommanderImpl.ListPaginated:", v)
		}

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			text,
		)

		if markup != nil {
			msg.ReplyMarkup = *markup
		}

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Return1CommanderImpl.ListPaginated: error sending reply message to chat [%v]", err)
		}
	}

	return1Elements, err := c.service.List(cursor, limit)
	if err != nil && err != return1.LastPageExceededErr {
		reply("error during getting list from service", nil, err)

		return
	}

	textResponse := "Here the elements:\n\n"
	for _, r := range return1Elements {
		textResponse += r.String() + "\n"
	}

	buttons := []tgbotapi.InlineKeyboardButton{}

	var prevButton *tgbotapi.InlineKeyboardButton

	if err != return1.LastPageExceededErr {
		nextButton := generateNextButton(cursor)
		if nextButton != nil {
			buttons = append(buttons, *nextButton)
		}
	}

	prevButton = generatePrevButton(cursor)

	if prevButton != nil {
		buttons = append(buttons, *prevButton)
	}

	msgMarkup := tgbotapi.NewInlineKeyboardMarkup(buttons)

	reply(textResponse, &msgMarkup)
}

func generateNextButton(cursor uint64) *tgbotapi.InlineKeyboardButton {
	paginationMarkupNext, _ := json.Marshal(PaginationMarkup{
		Cursor: cursor + PageSize,
		Limit:  PageSize,
	})

	cbPaginationNext := path.CallbackPath{
		Domain:       "exchange",
		Subdomain:    "return1",
		CallbackName: "list",
		CallbackData: string(paginationMarkupNext),
	}

	cbStringData := cbPaginationNext.String()

	return &tgbotapi.InlineKeyboardButton{
		Text:         "Next page",
		CallbackData: &cbStringData,
	}
}

func generatePrevButton(cursor uint64) *tgbotapi.InlineKeyboardButton {
	var prevPageCursor = int(cursor) - PageSize
	if prevPageCursor >= 0 {
		paginationMarkupPrev, _ := json.Marshal(PaginationMarkup{
			Cursor: uint64(prevPageCursor),
			Limit:  PageSize,
		})

		cbPaginationPrev := path.CallbackPath{
			Domain:       "exchange",
			Subdomain:    "return1",
			CallbackName: "list",
			CallbackData: string(paginationMarkupPrev),
		}

		cbStringData := cbPaginationPrev.String()

		return &tgbotapi.InlineKeyboardButton{
			Text:         "Prev page",
			CallbackData: &cbStringData,
		}
	}

	return nil
}
