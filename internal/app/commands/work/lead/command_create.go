package lead

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/work/lead"
	"log"
)

type createInputData struct {
	FirstName *string `json:"fName"`
	LastName  *string `json:"lName"`
	Project   *string `json:"prj"`
}

func (d *createInputData) valid() bool {
	return d.FirstName != nil && d.LastName != nil && d.Project != nil
}

func (c *LeadCommander) New(inputMessage *tgbotapi.Message) {
	parsedData := createInputData{}
	err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &parsedData)
	if err != nil || !parsedData.valid() {
		log.Printf("Work.LeadCommander.Create: "+
			"error reading json data for type createInputData from "+
			"input string %v - %v", inputMessage.CommandArguments(), err)
		c.sendMsg(inputMessage.Chat.ID, `Command format: /new__work__lead {"fName": "fName", "lName": "lName", "prj": "project"}`)
		return
	}

	leadId, err := c.leadService.Create(lead.Lead{
		FirstName: *parsedData.FirstName,
		LastName:  *parsedData.LastName,
		Project:   *parsedData.Project,
	})

	if err != nil {
		log.Printf("Work.LeadCommander.Create: %v", err)
		c.sendMsg(inputMessage.Chat.ID, fmt.Sprintf(`Error creating lead: %v`, err))
		return
	}

	c.sendMsg(inputMessage.Chat.ID, fmt.Sprintf(`Successfully created: /get__work__lead %d`, leadId))
}
