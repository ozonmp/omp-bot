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

	cursor, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Printf("BankOperationCommander.List: error converting argument 0 - %v", err)
		return
	}

	limit, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		log.Printf("BankOperationCommander.List: error converting argument 1 - %v", err)
		return
	}

	msg, err := c.prepareList(*inputMsg, cursor, limit)

	if err != nil {
		log.Printf("BankOperationCommander.List: error preparing list - %v", err)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankOperationCommander.List: error sending reply message to chat - %v", err)
	}
}

func(c *BankOperationCommander)  prepareList(inputMsg tgbotapi.Message, cursor uint64, limit uint64) (msg tgbotapi.MessageConfig, err error) {
	outputMsgText := "Here are the operations: \n\n"
	operations, err := c.operationService.List(cursor, limit)

	if err != nil {
		log.Printf("BankOperationCommander.List: error getting list - %v", err)
		return tgbotapi.MessageConfig{}, err
	}

	if len(operations) == 0 {
		msg = tgbotapi.NewMessage(inputMsg.Chat.ID, "No more operations")
		return msg, nil
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