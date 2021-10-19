package common

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c DummyCommonCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	commonId, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	_, err = c.commonService.Remove(uint64(commonId))
	var text string
	if err == nil {
		text = fmt.Sprintf("Delivery with id %d was deleted", commonId)
	} else {
		text = fmt.Sprintf("Fail to delete delivery: %v", err)
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyCommonCommander.Delete: error sending reply message to chat - %v", err)
	}

}
