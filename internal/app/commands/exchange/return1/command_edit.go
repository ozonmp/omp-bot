package return1

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/exchange"
)

func (c *Return1CommanderImpl) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	argsParsed := editCommandArgs{}

	err := json.Unmarshal([]byte(args), &argsParsed)
	if err != nil {
		replyToUser("error wrong input format "+
			showEditCommandInputFormat(),
			inputMsg,
			c.bot,
			err,
		)

		return
	}

	if argsParsed.Return1 == nil {
		replyToUser("error wrong input format. you forgot return1 "+
			showEditCommandInputFormat(),
			inputMsg,
			c.bot,
			err,
		)

		return
	}

	if argsParsed.ID == nil {
		replyToUser("error wrong input format. you forgot ID "+
			showEditCommandInputFormat(),
			inputMsg,
			c.bot,
			err,
		)

		return
	}

	err = c.service.Update(*argsParsed.ID, *argsParsed.Return1)
	if err != nil {
		replyToUser(fmt.Sprintf("failed edit element with idx[%d]", *argsParsed.ID), inputMsg, c.bot, err)

		return
	}

	replyToUser(fmt.Sprintf("Successfully edited element with id[%d]", *argsParsed.ID), inputMsg, c.bot)
}

type editCommandArgs struct {
	Return1 *exchange.Return1
	ID      *uint64
}

func showEditCommandInputFormat() string {
	return `[format should be { "Return1" : {  "Name" : "testName" }, "ID" : 0 } ]`
}
