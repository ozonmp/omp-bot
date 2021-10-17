package subdomain

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummySubdomainCommander) List(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	cursor, limit, err := parseListCmdArgs(args)
	if err != nil {
		log.Println("DummySubdomainCommander.List invalid args", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageList))
		return
	}
	pagination := NewPaginationList(c.subdomainService.List, cursor, limit)

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
		log.Printf("DummySubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}

func parseListCmdArgs(args string) (cursor, limit uint64, err error) {
	l := strings.Split(args, " ")
	if len(l) != 2 {
		return cursor, limit, errors.New(UsageList)
	}
	tmpCursor, err := strconv.Atoi(l[0])
	if tmpCursor <= 0 {
		return 0, 0, fmt.Errorf("invalid arg")
	} else if err != nil {
		return 0, 0, err
	}

	tmpLimit, err := strconv.Atoi(l[1])
	if tmpLimit <= 0 {
		return 0, 0, fmt.Errorf("invalid arg")
	} else if err != nil {
		return 0, 0, err
	}
	cursor = uint64(tmpCursor)
	limit = uint64(tmpLimit)
	return cursor, limit, nil
}
