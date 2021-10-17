package subdomain

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummySubdomainCommander) Help(inputMsg *tgbotapi.Message) {
	txtCmds := []string{
		fmt.Sprintf("/%s — print list of commands", CommandHelp),
		fmt.Sprintf("/%s — get an entity", CommandGet),
		fmt.Sprintf("/%s — get a list of your entity", CommandList),
		fmt.Sprintf("/%s — delete an existing entity", CommandDelete),
		fmt.Sprintf("/%s — create a new entity", CommandNew),
		fmt.Sprintf("/%s — edit an entity", CommandEdit),
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, strings.Join(txtCmds, "\n"))
	c.bot.Send(msg)
}
