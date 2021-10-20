package transaction

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BankTransactionCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the transactions: \n\n"

	transactions, err := c.transactionService.List(c.cursor, c.limit)
	if err != nil {
		outputMsgText = err.Error()
	} else {
		for _, p := range transactions {
			outputMsgText += p.String()
			outputMsgText += "\n"
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	if c.cursor > 0 {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardButtonData("Prev", "bank__transaction__list__prev"),
		)
	}

	if uint64(len(transactions)) == c.limit {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardButtonData("Next", "bank__transaction__list__next"),
		)
	}

	if len(buttons) > 0 {
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
		msg.ReplyMarkup = keyboard
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankTransactionCommander.List: error sending reply message to chat - %v", err)
	}
}
