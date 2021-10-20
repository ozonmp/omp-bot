package singlesubscription

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummySingleSubscriptionCommander) List(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	arg, err := parseListCmdArgs(args)
	if err != nil {
		log.Println("DummySingleSubscriptionCommander.List invalid args", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageList))
		return
	}
	pagination := NewPaginationList(c.service.List, arg.Cursor, arg.Limit)

	page := pagination.Page()
	buttons := pagination.Buttons()

	inlineKeyboardButtons := make([]tgbotapi.InlineKeyboardButton, 0, 2)
	for _, v := range buttons {
		inlineKeyboardButtons = append(
			inlineKeyboardButtons,
			tgbotapi.NewInlineKeyboardButtonData(v.Text, v.Data),
		)
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, page)
	if len(inlineKeyboardButtons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(inlineKeyboardButtons...),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.List: error sending reply message to chat - %v", err)
	}
}

func parseListCmdArgs(args string) (*CallbackListData, error) {
	l := strings.Split(args, " ")
	data := &CallbackListData{}
	if len(l) != 2 {
		return data, errors.New(UsageList)
	}
	tmpCursor, err := strconv.Atoi(l[0])
	if err != nil {
		return data, err
	} else if tmpCursor <= 0 {
		return data, fmt.Errorf("invalid arg")
	}

	tmpLimit, err := strconv.Atoi(l[1])
	if err != nil {
		return data, err
	} else if tmpLimit <= 0 {
		return data, fmt.Errorf("invalid arg")
	}

	data.Cursor = uint64(tmpCursor)
	data.Limit = uint64(tmpLimit)

	return data, nil
}
