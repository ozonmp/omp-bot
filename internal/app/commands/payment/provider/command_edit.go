package provider

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	provModels "github.com/ozonmp/omp-bot/internal/service/payment/provider"
	"log"
)

func (c *PaymentProviderCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var p provModels.Provider
	err := json.Unmarshal([]byte(args), &p)

	if err != nil {
		log.Printf("Invalid json.\n%v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Got invalid json string!"))
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("PaymentProviderCommander.Edit: error sending reply message to chat - %v", err)
		}
		return
	}

	_, err = c.providerService.Get(p.Id)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%v", err))
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("PaymentProviderCommander.Create: error sending reply message to chat - %v", err)
		}
		return
	}

	err = c.providerService.Update(p.Id, p)
	var outputText string
	if err != nil {
		log.Printf("failed to update provider %v", err)
		outputText = "Failed to update provider"
	} else {
		outputText = fmt.Sprintf("Provider with ID:%d has been updated.", p.Id)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PaymentProviderCommander.Edit: error sending reply message to chat - %v", err)
	}
}
