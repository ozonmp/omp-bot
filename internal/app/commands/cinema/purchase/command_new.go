package purchase

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

func (c *PurchaseCommanderImpl) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	if len(args) == 0 {
		replyToUser("some error. you should provide name for element", inputMsg, c.bot)

		return
	}

	parsedPurchase := cinema.Purchase{}

	err := json.Unmarshal([]byte(args), &parsedPurchase)
	if err != nil {
		replyToUser(
			"error creating new element. wrong format "+
				cinema.ShowPurchaseInputFormat(),
			inputMsg,
			c.bot,
			err,
		)
	}

	if len(parsedPurchase.Name) == 0 {
		replyToUser("some error. you should provide name for element", inputMsg, c.bot)

		return
	}

	id, err := c.service.Create(parsedPurchase)
	if err != nil {
		replyToUser("error creating new element", inputMsg, c.bot, err)

		return
	}

	replyToUser(fmt.Sprintf("New element created with id[%d]", id), inputMsg, c.bot)
}
