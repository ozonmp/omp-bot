package rent

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CinemaRentCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData, err := c.jsonListCallbackParser(callbackPath)
	if err != nil {
		log.Printf("CinemaRentCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	c.sendList(callback.Message.Chat.ID, uint64(parsedData.Cursor), uint64(parsedData.Limit))

	// msg := tgbotapi.NewMessage(
	// 	callback.Message.Chat.ID,
	// 	fmt.Sprintf("Parsed: %+v\n", parsedData),
	// )

	// _, err = c.bot.Send(msg)
	// if err != nil {
	// 	log.Printf("CinemaRentCommander.CallbackList: error sending reply message to chat - %v", err)
	// }
}
