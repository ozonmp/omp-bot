package group

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *ProductGroupCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText := ""
	for id := range c.groups {
		outputMsgText += fmt.Sprintf("%d. %s\n", id, c.groups[id].String(), c.groups[id].String())
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	c.bot.Send(msg)
}
