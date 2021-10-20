package warehouse

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EstateWarehouseCommander) Describe(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		str := "Command format: /get__estate__warehouse $WAREHOUSE_ID"
		log.Println("Wrong arguments: '", args, "'. %s", str)
		c.sendMsg(inputMessage.Chat.ID, str)
		return
	}

	warehouse, err := c.warehouseService.Describe(uint64(idx))

	if err != nil {
		str := fmt.Sprintf("Failed to get warehouse with id %d: %v", idx, err)
		log.Printf(str)
		c.sendMsg(inputMessage.Chat.ID, str)
		return
	}

	str_message := strconv.FormatUint(warehouse.ID, 10) + " " +
		strconv.FormatUint(warehouse.OwnerID, 10) + " " + warehouse.Address

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		str_message,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EstateWarehouseCommander.Get: error sending reply message to chat - %v", err)
	}
}
