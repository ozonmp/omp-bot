package workplace

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
)

var WorkplaceJsonExample string

func init() {
	workplaceExample, _ := json.Marshal(business.Workplace{ID: 1, Title: "Title Example", EmployeeID: 1, WorkplaceNumber: 1, OfficeID: 1})
	WorkplaceJsonExample = string(workplaceExample)
}

func (c *BusinessWorkplaceCommander) processError(tgChatID int64, description string, solution string) {
	log.Println(description)
	var msg = tgbotapi.NewMessage(tgChatID, description+"\n"+solution)
	c.bot.Send(msg)
}
