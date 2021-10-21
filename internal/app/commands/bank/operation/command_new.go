package operation

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/bank"
	"log"
	"strconv"
	"strings"
)

func (c BankOperationCommander) New(inputMsg *tgbotapi.Message){
	args := strings.Split(inputMsg.CommandArguments(), " ")
	if len(args) != 2 {
		log.Printf("BankOperationCommander.New: arguments amount not 2, but - %v", len(args))
		return
	}

	operationType := bank.OperationType(args[0])

	transactionID, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		log.Printf("BankOperationCommander.New: error converting argument 2 - %v", err)
		return
	}

	op, err := c.operationService.Create(bank.NewOperation(operationType, transactionID))
	if err != nil {
		log.Printf("fail to create operation: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("OperationID - %d", op),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("OperationCommander.New: error sending reply message to chat - %v", err)
	}
}
