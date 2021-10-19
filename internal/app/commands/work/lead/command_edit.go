package lead

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/work/lead"
	"log"
)

type editInputData struct {
	ID        *uint64 `json:"id"`
	FirstName *string `json:"fName"`
	LastName  *string `json:"lName"`
	Project   *string `json:"prj"`
}

func (d *editInputData) valid() bool {
	return d.FirstName != nil && d.LastName != nil && d.Project != nil && d.ID != nil
}

func (c *LeadCommander) Edit(inputMessage *tgbotapi.Message) {
	parsedData := editInputData{}
	err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &parsedData)
	if err != nil || !parsedData.valid() {
		log.Printf("Work.LeadCommander.Edit: "+
			"error reading json data for type editInputData from "+
			"input string %v - %v", inputMessage.CommandArguments(), err)
		c.sendMsg(inputMessage.Chat.ID, `Command format: /edit__work__lead {"id": 1, "fName": "fName", "lName": "lName", "prj": "project"}`)
		return
	}

	err = c.leadService.Update(*parsedData.ID, lead.Lead{
		FirstName: *parsedData.FirstName,
		LastName:  *parsedData.LastName,
		Project:   *parsedData.Project,
	})

	if err != nil {
		log.Printf("Work.LeadCommander.Edit: %v", err)
		c.sendMsg(inputMessage.Chat.ID, fmt.Sprintf(`Error updating lead: %v`, err))
		return
	}

	c.sendMsg(inputMessage.Chat.ID, fmt.Sprintf(`Successfully updated: /get__work__lead %d`, *parsedData.ID))
}
