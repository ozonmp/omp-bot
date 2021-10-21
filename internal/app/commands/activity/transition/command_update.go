package transition

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/activity"
)

func (c *ActivityTransitionCommander) Update(inputMessage *tgbotapi.Message) {
	var id uint64
	var err error

	parsedData := CreateTransition{}

	argsS := inputMessage.CommandArguments()

	args := strings.Fields(argsS)

	if len(args) == 2 {
		id, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Println("wrong args", args[0])
			return
		}

		err := json.Unmarshal([]byte(args[1]), &parsedData)
		if err != nil {
			log.Printf("ActivityTransitionCommander.Update: "+
				"error reading json data for type CreateTransition from "+
				"input string %v - %v", args, err)
			return
		}
	}

	newTransition := activity.Transition{
		Id:   id,
		Name: parsedData.Name,
		From: parsedData.From,
		To:   parsedData.To,
	}

	var txtResult string
	err = c.transitionService.Update(id, newTransition)
	if err != nil {
		log.Printf("Fail to update transition with id: %d. %v", id, err)
		txtResult = "Failed"
	} else {
		txtResult = "Successful"
	}

	textMsg := fmt.Sprintf("%s update transition", txtResult)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		textMsg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityTransitionCommander.Update: error sending reply message to chat - %v", err)
	}
}
