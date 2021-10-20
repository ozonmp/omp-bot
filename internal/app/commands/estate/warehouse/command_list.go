package warehouse

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const maxElemInResponse = 4

func (c *EstateWarehouseCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the warehouses: \n\n"
	log.Printf("EstateWarehouseCommander.List: obtain list of warehouses")

	warehouses, err := c.warehouseService.List(0, maxElemInResponse)
	if err != nil {
		log.Printf("EstateWarehouseCommander.List: obtain list of warehouses- %v", err)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
		_, err = c.bot.Send(msg)
		return
	}

	if len(warehouses) == 0 {
		outputMsgText += "No items\n"
	}

	for _, p := range warehouses {
		outputMsgText += strconv.FormatUint(p.ID, 10) + " " +
			strconv.FormatUint(p.OwnerID, 10) + " " +
			p.Address + " " +
			strconv.FormatUint(uint64(p.AreaM2), 10) + " " +
			strconv.FormatUint(p.PriceInCents, 10) +
			"\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("EstateWarehouseCommander.List: error sending reply message to chat - %v", err)
	}
}
