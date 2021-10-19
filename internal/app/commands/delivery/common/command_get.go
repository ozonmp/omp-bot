package common

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c DummyCommonCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	commonId, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong args", args)
		return
	}

	common, err := c.commonService.Describe(uint64(commonId))
	var text string
	if err == nil {
		text = common.String()
	} else {
		text = fmt.Sprintf("Fail to get delivery: %v", err)
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyCommonCommander.Get: error sending reply message to chat - %v", err)
	}
}
