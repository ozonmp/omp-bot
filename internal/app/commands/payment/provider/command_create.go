package provider

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	provModels "github.com/ozonmp/omp-bot/internal/service/payment/provider"
	"log"
	"strings"
)

func (c *PaymentProviderCommander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var p provModels.Provider
	err := json.Unmarshal([]byte(args), &p)

	if err != nil {
		log.Printf("Invalid json.\n%v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Got invalid json string!"))
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("PaymentProviderCommander.Create: error sending reply message to chat - %v", err)
		}
		return
	}

	if len(strings.TrimSpace(p.Name)) == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Name field is required!")
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("PaymentProviderCommander.Create: error sending reply message to chat - %v", err)
		}
		return
	}

	id, err := c.providerService.Create(p)
	var outputText string
	if err != nil {
		log.Printf("failed to create provider %v", err)
		outputText = "Failed to create provider"
	} else {
		outputText = fmt.Sprintf("New payment provider created. ID:%d", id)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PaymentProviderCommander.Get: error sending reply message to chat - %v", err)
	}
}
