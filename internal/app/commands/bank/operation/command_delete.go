package operation

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *BankOperationCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	ok, err := c.operationService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to delete operation with idx %d: %v", idx, err)
		return
	}

	var msgText string
	if ok {
		msgText = fmt.Sprintf("operation with ID %v has deleted", idx)
	} else {
		msgText = fmt.Sprintf("operation with ID %v has not deleted", idx)
	}
	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		msgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("OperationCommander.Delete: error sending reply message to chat - %v", err)
	}
}
