package production

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const textSuccessDelete = "successfully deleted"

func (c *RecommendationProductionCommander) Delete(inputMessage *tgbotapi.Message) {
	text, err := func() (string, error) {
		args := inputMessage.CommandArguments()

		id, err := strconv.ParseUint(args, 0, 64)
		if err != nil {
			return textWrong, fmt.Errorf("%w %v %v", errWrongArgs, args, err)
		}

		err = c.productionService.Delete(id)
		if err != nil {
			return textWrong, err
		}

		return textSuccessDelete, nil
	}()
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	c.sendMessage(msg)
}
