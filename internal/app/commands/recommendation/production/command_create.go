package production

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/recommendation"
)

const textSuccessCreate = "successfully created"

func (c *RecommendationProductionCommander) Create(inputMessage *tgbotapi.Message) {
	text, err := func() (string, error) {
		var production recommendation.Production

		args := inputMessage.CommandArguments()

		err := json.Unmarshal([]byte(args), &production)
		if err != nil {
			return textWrong, fmt.Errorf("%w %v", errWrongArgs, args)
		}

		err = c.productionService.Create(production)
		if err != nil {
			return textWrong, err
		}

		return textSuccessCreate, nil
	}()
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	c.sendMessage(msg)
}
