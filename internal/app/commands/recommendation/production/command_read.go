package production

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RecommendationProductionCommander) Read(inputMessage *tgbotapi.Message) {
	text, err := func() (string, error) {
		args := inputMessage.CommandArguments()

		id, err := strconv.ParseUint(args, 0, 64)
		if err != nil {
			return textWrong, fmt.Errorf("%w %v %v", errWrongArgs, args, err)
		}

		production, err := c.productionService.Read(id)
		if err != nil {
			return textWrong, err
		}

		return production.String(), nil
	}()
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	c.sendMessage(msg)
}
