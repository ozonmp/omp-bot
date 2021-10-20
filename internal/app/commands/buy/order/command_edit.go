package order

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/buy"
)

type EditOrderData struct {
	Id       uint64 `json:"id"`
	Title    string `json:"title"`
	Quantity uint64 `json:"quantity"`
}

func (c *OrderCommander) Edit(inputMessage *tgbotapi.Message) {
	data := inputMessage.CommandArguments()

	parsedData := EditOrderData{}
	err := json.Unmarshal([]byte(data), &parsedData)
	if err != nil {
		log.Printf("OrderCommander.Edit: "+
			"error reading json data from "+
			"input string %v - %v", data, err)

		c.Reply(
			inputMessage.Chat.ID,
			"Failed to parse json! Correct syntax for 'edit' command is:\n"+
				`/edit__buy__order {"id": <number>, "title": <string>, "quantity": <number>}`)
		return
	}

	err = c.orderService.Update(
		parsedData.Id,
		buy.Order{
			Title:    parsedData.Title,
			Quantity: parsedData.Quantity,
		})

	if err != nil {
		log.Printf("Fail to update order: %v", err)
		c.Reply(
			inputMessage.Chat.ID,
			fmt.Sprintf("Fail to update order: %v", err))
		return
	}

	c.Reply(
		inputMessage.Chat.ID,
		fmt.Sprintf("Order with id %d updated successfully", parsedData.Id))
}
