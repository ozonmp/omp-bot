package operation

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c BankOperationCommander) Get(inputMsg *tgbotapi.Message){
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	op, err := c.operationService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get operation with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		op.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("OperationCommander.Get: error sending reply message to chat - %v", err)
	}
}