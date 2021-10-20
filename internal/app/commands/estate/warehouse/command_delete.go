package warehouse

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EstateWarehouseCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	ID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("Error during parse id of warehouse in Delete(): %s", err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			`/delete__estate__warehouse $WAREHOUSE_ID.
			Expect WAREHOUSE_ID as non negative integer number`,
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("Error during sending error message to chat - %v", err)
		}
		return
	}

	was_deleted, err := c.warehouseService.Delete(ID)
	if err != nil {
		str := fmt.Sprintf("Failed to delete warehouse with id '%v': %s", ID, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			str,
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("EstateWarehouseCommander.Delete: error sending reply message to chat - %v", err)
		}
		return
	}

	var str_msg string
	if was_deleted {
		str_msg = fmt.Sprintf("Succesfuly delete warehouse with id %d", ID)
	} else {
		str_msg = fmt.Sprintf("Warehouse with id %d not found", ID)
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		str_msg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EstateWarehouseCommander.Delete: error sending reply message to chat - %v", err)
	}

	log.Printf(str_msg)
}
