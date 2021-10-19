package verification

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const maxElementsSize uint64 = 5

func (c *ServiceVerificationCommander) List(position uint64, inputMessage *tgbotapi.Message) {
	outputMsgText := "Items: \n\n"

	products, _ := c.verificationService.List(position, maxElementsSize)
	for _, p := range products {
		outputMsgText += fmt.Sprintf("%v\n", p.String())
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(c.createKeyboard(position, uint64(len(products)))...))


	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}

