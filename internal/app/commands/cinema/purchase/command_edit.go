package purchase

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

func (c *PurchaseCommanderImpl) Edit(inputMsg *tgbotapi.Message) {
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

	if argsParsed.Purchase == nil {
		replyToUser("error wrong input format. you forgot purchase "+
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

	err = c.service.Update(*argsParsed.ID, *argsParsed.Purchase)
	if err != nil {
		replyToUser(fmt.Sprintf("failed edit element with idx[%d]", *argsParsed.ID), inputMsg, c.bot, err)

		return
	}

	replyToUser(fmt.Sprintf("Successfully edited element with id[%d]", *argsParsed.ID), inputMsg, c.bot)
}

type editCommandArgs struct {
	Purchase *cinema.Purchase
	ID       *uint64
}

func showEditCommandInputFormat() string {
	return `[format should be { "Purchase" : {  "Name" : "testName" }, "ID" : *integer* } ]`
}
