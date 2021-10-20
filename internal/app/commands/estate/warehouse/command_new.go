package warehouse

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/estate/warehouse"
)

func (c *EstateWarehouseCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	if len(args) == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Warehouse records have to be provided.")
		c.bot.Send(msg)
		return
	}

	warehouseFromChat := warehouse.Warehouse{}
	if args[0] == '{' {
		err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &warehouseFromChat)
		if err != nil {
			str := fmt.Sprintf(`Failed to parse warehouse records: '%v'`, inputMessage.CommandArguments())
			log.Println(str)
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, str)
			c.bot.Send(msg)
			return
		}
	} else {
		str := fmt.Sprintf(`You must create new warehouse only via json-way.
							It's not valid json: '%v'`, inputMessage.CommandArguments())
		log.Println(str)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, str)
		c.bot.Send(msg)
		return
	}

	newId, err := c.warehouseService.New(warehouseFromChat)
	if err != nil {
		str := fmt.Sprintf("Error during creation new warehouse: %s", err.Error())
		log.Println(str)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, str)
		c.bot.Send(msg)
		return
	}

	str := fmt.Sprintf("New warehouse with id '%v' was created successfully.", newId)
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		str,
	)
	c.bot.Send(msg)
	log.Println(str)
}
