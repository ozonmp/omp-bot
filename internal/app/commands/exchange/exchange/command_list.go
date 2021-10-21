package exchange

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	cursorStep uint64 = 5
)

func (c *SubdomainCommander) List(inputMsg *tgbotapi.Message) {
	c.showPage(inputMsg, 0)
}

func (c *SubdomainCommander) showPage(inputMsg *tgbotapi.Message, cursor uint64) {
	outputMsgText := ""

	exchangeRequestList, err := c.exchangeService.List(cursor, cursorStep)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
			"You've reached end of list",
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("SubdomainCommander.List: error sending reply message to chat - %v", err)
		}
		return
	}
	for _, er := range exchangeRequestList {
		outputMsgText += fmt.Sprintf("ID: %v, Package: %v\nFrom: %v, To: %v\nStatus: %v\n",
			er.Id, er.Package, er.From, er.To, er.Status)
		outputMsgText += "\n"
	}

	var nextCursor, prevCursor uint64
	nextCursor = cursor + uint64(len(exchangeRequestList))
	if cursor == 0 {
		prevCursor = 0
	} else if cursor % cursorStep != 0 {
		prevCursor = cursor - (cursor % cursorStep)
	} else {
		prevCursor = cursor - cursorStep
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	var buttons []tgbotapi.InlineKeyboardButton

	exchangeRequestList, _ = c.exchangeService.List(cursor + 1, cursorStep)
	hasNext     := uint64(len(exchangeRequestList)) >= cursorStep
	hasPrevious := cursor > 0

	serializedDataNext, _ := json.Marshal(CallbackListData{
		Cursor: nextCursor,
	})

	serializedDataPrev, _ := json.Marshal(CallbackListData{
		Cursor: prevCursor,
	})

	callbackNextPath := path.CallbackPath{
		Domain:       "exchange",
		Subdomain:    "exchange",
		CallbackName: "list",
		CallbackData: string(serializedDataNext),
	}

	callbackPrevPath := path.CallbackPath{
		Domain:       "exchange",
		Subdomain:    "exchange",
		CallbackName: "list",
		CallbackData: string(serializedDataPrev),
	}

	if hasPrevious {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPrevPath.String()))
	}
	if hasNext {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next page", callbackNextPath.String()))
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttons...),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("SubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}