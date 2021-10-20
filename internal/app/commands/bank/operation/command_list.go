package operation

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *BankOperationCommander) List(inputMsg *tgbotapi.Message) {
	args := strings.Split(inputMsg.CommandArguments(), " ")
	if len(args) != 2 {
		log.Printf("BankOperationCommander.List: arguments amount not 2, but - %v", len(args))
		return
	}

	cursor, errArg := strconv.ParseUint(args[0], 10, 64)
	if errArg != nil {
		log.Printf("BankOperationCommander.List: error converting argument 0 - %v", errArg)
		return
	}

	limit, errArg := strconv.ParseUint(args[1], 10, 64)
	if errArg != nil {
		log.Printf("BankOperationCommander.List: error converting argument 1 - %v", errArg)
		return
	}

	msg, errList := c.prepareList(*inputMsg, cursor, limit)

	if errList != nil {
		log.Printf("BankOperationCommander.List: error preparing list - %v", errList)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankOperationCommander.List: error sending reply message to chat - %v", err)
	}
}

func(c *BankOperationCommander)  prepareList(inputMsg tgbotapi.Message, cursor uint64, limit uint64) (msg tgbotapi.MessageConfig, err error) {
	outputMsgText := "Here are the operations: \n\n"
	operations, errList := c.operationService.List(cursor, limit)

	if len(operations) == 0 {
		msg = tgbotapi.NewMessage(inputMsg.Chat.ID, "No more operations")
		return msg, nil
	}

	if errList != nil {
		log.Printf("BankOperationCommander.List: error getting list - %v", errList)
		return tgbotapi.MessageConfig{}, errList
	}
	for _, op := range operations {
		outputMsgText += op.String()
		outputMsgText += "\n\n"
	}

	msg = tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: cursor + limit,
		Limit: limit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "bank",
		Subdomain:    "operation",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	return msg, nil
}