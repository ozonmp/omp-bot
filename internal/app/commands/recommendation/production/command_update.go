package production

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/recommendation"
)

const (
	textSuccessUpdate = "successfully updated"
)

func (c *RecommendationProductionCommander) Update(inputMessage *tgbotapi.Message) {
	text, err := func() (string, error) {
		var production recommendation.Production

		args := inputMessage.CommandArguments()

		err := json.Unmarshal([]byte(args), &production)
		if err != nil {
			return textWrong, fmt.Errorf("%w %v", errWrongArgs, args)
		}

		err = c.productionService.Update(production)
		if err != nil {
			return textWrong, err
		}

		return textSuccessUpdate, nil
	}()
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	c.sendMessage(msg)
}
