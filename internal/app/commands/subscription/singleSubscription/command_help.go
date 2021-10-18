package singleSubscription

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummySingleSubscriptionCommander) Help(inputMsg *tgbotapi.Message) {
	txtCmds := []string{
		fmt.Sprintf("/%s — print list of commands", CommandHelp),
		fmt.Sprintf("/%s — get an entity", CommandGet),
		fmt.Sprintf("/%s — get a list of your entity", CommandList),
		fmt.Sprintf("/%s — delete an existing entity", CommandDelete),
		fmt.Sprintf("/%s — create a new entity", CommandNew),
		fmt.Sprintf("/%s — edit an entity", CommandEdit),
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, strings.Join(txtCmds, "\n"))

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.Help: error sending reply message to chat - %v", err)
	}
}
