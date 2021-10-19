package order

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/buy"
)

type NewOrderData struct {
	Title    string `json:"title"`
	Quantity uint64 `json:"quantity"`
}

func (c *OrderCommander) New(inputMessage *tgbotapi.Message) {
	data := inputMessage.CommandArguments()

	parsedData := NewOrderData{}
	err := json.Unmarshal([]byte(data), &parsedData)
	if err != nil {
		log.Printf("OrderCommander.New: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", data, err)

		c.Reply(
			inputMessage.Chat.ID,
			"Failed to parse json! Correct syntax for 'new' command is:\n"+
				`/new__buy__order {"title": <string>, "quantity": <number>}`)
		return
	}

	newId, err := c.orderService.Create(
		buy.Order{
			Title:    parsedData.Title,
			Quantity: parsedData.Quantity,
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
