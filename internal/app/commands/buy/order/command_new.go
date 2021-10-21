package order

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/buy"
)

type NewOrderData struct {
	UserId    uint64 `json:"user_id"`
	AddressId uint64 `json:"address_id"`
	StateId   uint32 `json:"state_id"`
	Paid      bool   `json:"paid"`
}

func (c *OrderCommander) New(inputMessage *tgbotapi.Message) {
	data := inputMessage.CommandArguments()

	parsedData := NewOrderData{}
	err := json.Unmarshal([]byte(data), &parsedData)
	if err != nil {
		log.Printf("OrderCommander.New: "+
			"error reading json data from "+
			"input string %v - %v", data, err)

		c.Reply(
			inputMessage.Chat.ID,
			"Failed to parse json! Correct syntax for 'new' command is:\n"+
				`/new__buy__order {"user_id": <number>, "address_id": <number>, "state_id":<number>, "paid": <true|false>}`)
		return
	}

	newId, err := c.orderService.Create(
		buy.Order{
			UserId:    parsedData.UserId,
			AddressId: parsedData.AddressId,
			StateId:   parsedData.StateId,
			Paid:      parsedData.Paid,
		})

	if err != nil {
		log.Printf("Fail to create order: %v", err)
		return
	}

	c.Reply(
		inputMessage.Chat.ID,
		fmt.Sprintf("Order with id %v created", newId),
	)
}
