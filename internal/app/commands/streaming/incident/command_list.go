package incident

import (
	"encoding/json"
	"errors"
	"log"
	"sort"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

var EOF = errors.New("end of list")

const pageSize = 3

func (c *Commander) List(inputMessage *tgbotapi.Message) {

	msg, _ := c.prepareMsgList(inputMessage.Chat.ID, 0, pageSize)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: pageSize,
	})

	callbackPath := path.CallbackPath{
		Domain:       "streaming",
		Subdomain:    "incident",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	c.send(msg)

}

func (c *Commander) send(msg tgbotapi.MessageConfig) {
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingIncidentCommander.List: error sending reply message to chat - %v", err)
	}
}

func (c *Commander) prepareMsgList(chatId int64, from, to int) (tgbotapi.MessageConfig, error) {
	outputMsgText := "Here current page of the incidents: \n\n"

	incidents := c.incidentService.List()
	sort.SliceStable(incidents, func(i, j int) bool {
		return incidents[i].Id < incidents[j].Id
	})
	if from >= len(incidents) {
		return tgbotapi.NewMessage(chatId, "next page out of range"), errors.New("next page out of range")
	}
	var err error
	if to >= len(incidents) {
		to = len(incidents)
		err = EOF
	}

	for _, entity := range incidents[from:to] {
		rawEntity, err := entity.String()
		if err != nil {
			return tgbotapi.NewMessage(chatId, "internal server error"), err
		}
		outputMsgText += rawEntity
		outputMsgText += "\n"
	}

	return tgbotapi.NewMessage(chatId, outputMsgText), err
}
