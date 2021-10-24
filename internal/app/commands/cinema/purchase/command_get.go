package purchase

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PurchaseCommanderImpl) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		replyToUser("error wrong arguments. should be 1 number - element id.", inputMsg, c.bot, err)

		return
	}

	element, err := c.service.Describe(uint64(idx))
	if err != nil {
		replyToUser(fmt.Sprintf("failed get with idx[%d]", idx), inputMsg, c.bot, err)

		return
	}

	replyToUser("Successfully got element: "+element.String(), inputMsg, c.bot)
}
