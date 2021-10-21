package transition

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/activity"
)

func (c *ActivityTransitionCommander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	parsedData := CreateTransition{}
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		log.Printf("ActivityTransitionCommander.Create: "+
			"error reading json data for type CreateTransition from "+
			"input string %v - %v", args, err)
		return
	}

	newTransition := activity.Transition{
		Id:   0,
		Name: parsedData.Name,
		From: parsedData.From,
		To:   parsedData.To,
	}

	var txtResult string
	_, err = c.transitionService.Create(newTransition)
	if err != nil {
		log.Printf("Fail to create transition: %v", err)
		txtResult = "Failed"
	} else {
		txtResult = "Successful"
	}

	textMsg := fmt.Sprintf("%s create transition", txtResult)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		textMsg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityTransitionCommander.Delete: error sending reply message to chat - %v", err)
	}
}
