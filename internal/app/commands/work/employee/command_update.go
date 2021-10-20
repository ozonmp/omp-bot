package employee

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (c *DemoSubdomainCommander) Update(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.CommandArguments(), " ")

	if len(args) != 2 {
		c.sendWrongArgs(inputMessage)
		return
	}

	employeeId, err := strconv.Atoi(args[0])
	if err != nil {
		c.sendWrongArgs(inputMessage)
		return
	}

	c.sendMessage(inputMessage, c.subdomainService.Update(employeeId, args[1]))
}
