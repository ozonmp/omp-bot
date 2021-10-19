package ground

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
)

type groundRawData struct {
	Name        string `json:"name"`
	WheelsCount uint64 `json:"wheels"`
	Color       string `json:"color"`
	MaxSpeed    uint64 `json:"speed"`
}

func (c *GroundCommander) New(inputMessage *tgbotapi.Message) {
	var msgText string

	defer func() {
		c.Send(
			inputMessage.Chat.ID,
			msgText,
		)
	}()

	args := inputMessage.CommandArguments()
	var rawData groundRawData

	err := json.Unmarshal([]byte(args), &rawData)

	if err != nil {
		log.Printf("Internal error %v", err)
		msgText = `Format command is bad.
Format: {"name":"{ground_name}", "wheels":{wheels_count}, "color":"{color}, "speed":{max_speed}}`
		return
	}

	idx, err := c.service.Create(
		autotransport.Ground{
			Name:        rawData.Name,
			WheelsCount: rawData.WheelsCount,
			Color:       rawData.Color,
			MaxSpeed:    rawData.MaxSpeed,
		},
	)
	if err == nil {
		msgText = fmt.Sprintf("Success. GroundID = %d", idx)
	} else {
		log.Printf("Internal error %v", err)
		msgText = err.Error()
	}
}
