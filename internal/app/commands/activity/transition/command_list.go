package transition

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	Offset uint64 = 0
	Limit  uint64 = 2
)

func internalList(c *ActivityTransitionCommander, chatId int64, offset uint64, limit uint64) {
	outputMsgText := fmt.Sprintf("Here list of transitions from %v to %v: \n\n", offset, offset+limit)

	transitions, _ := c.transitionService.List(offset, limit)
	for _, p := range transitions {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(chatId, outputMsgText)

	if offset+limit < c.transitionService.Size() {
		serializedData, _ := json.Marshal(CallbackListData{
			Offset: offset + limit,
			Limit:  limit,
		})

		callbackPath := path.CallbackPath{
			Domain:       "activity",
			Subdomain:    "transition",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityTransitionCommander.List: error sending reply message to chat - %v", err)
	}
}

func (c *ActivityTransitionCommander) List(inputMessage *tgbotapi.Message) {
	var offset, limit = Offset, Limit
	var err error

	argsS := inputMessage.CommandArguments()

	args := strings.Fields(argsS)

	if len(args) == 2 {
		offset, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Println("wrong args", args[0])
			return
		}
		limit, err = strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			log.Println("wrong args", args[1])
			return
		}
	}

	internalList(c, inputMessage.Chat.ID, offset, limit)
}
