package return1

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *Return1CommanderImpl) ListCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := PaginationMarkup{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("Return1CommanderImpl.ListCallback: "+
			"error reading json data for type PaginationMarkup from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	c.ListPaginated(callback.Message, parsedData.Cursor, parsedData.Limit)
}
