package warehouse

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/estate/warehouse"
)

func (c *EstateWarehouseCommander) Edit(inputMessage *tgbotapi.Message) {
	splitArgs := strings.SplitN(inputMessage.CommandArguments(), ", ", 2)
	if len(splitArgs) != 2 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
			fmt.Sprintf("WarehousID and new json-records have to be provided"))
		c.bot.Send(msg)
		return
	}

	warehouseID, err := strconv.ParseUint(splitArgs[0], 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Failed to parse warehouse ID")
		c.bot.Send(msg)
		return
	}

	warehouseFromChat := warehouse.Warehouse{}
	err = json.Unmarshal([]byte(splitArgs[1]), &warehouseFromChat)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
			fmt.Sprintf(`Failed to parse warehouse "%v"`, splitArgs[1]))
		c.bot.Send(msg)
		return
	}

	err = c.warehouseService.Edit(warehouseID, warehouseFromChat)
	if err != nil {
		str := fmt.Sprintf("Failed to edit warehouse with id '%v': %s", warehouseID, err.Error())
		log.Printf(str)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, str)
		c.bot.Send(msg)
	}

	str := fmt.Sprintf("Warehouse with id '%v' was edited successfully.", warehouseID)
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		str,
	)
	c.bot.Send(msg)
	log.Printf(str)
}
