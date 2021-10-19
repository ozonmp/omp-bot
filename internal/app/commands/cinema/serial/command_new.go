package serial

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/cinema/serial"
)

func (c *CinemaSerialCommander) New(inputMessage *tgbotapi.Message) {
	msg := ""

	defer func() {
		tgmsg := tgbotapi.NewMessage(inputMessage.Chat.ID, msg)
		_, err := c.bot.Send(tgmsg)
		if err != nil {
			log.Printf("CinemaSerialCommander.Get: error sending reply message to chat (%v)", err)
		}
	}()

	args := inputMessage.CommandArguments()
	argss := strings.Split(args, "\n")

	// fmt.Println(args, inputMessage.Text, argss)

	idx, err := strconv.Atoi(argss[0])
	if err != nil || idx == 0 {
		msg = "incorrect id=" + argss[0]
		log.Println(msg)
		return
	}

	if len(argss[1]) == 0 || len(argss[1]) > 200 {
		msg = "incorrect title=" + argss[1]
		log.Println(msg)
		return
	}
	// TODO validate string
	title := argss[1]

	year, err := strconv.Atoi(argss[2])
	if err != nil || year < 1800 || year > 2021 {
		//int64(time.Now().Year) {
		msg = "incorrect year=" + argss[2]
		log.Println(msg)
		return
	}

	// newserial := Serial{ID}
	var new serial.Serial
	new.ID = idx
	new.Title = title
	new.Year = year

	err = c.subdomainService.New(new)
	if err != nil {
		msg = fmt.Sprintf("Err: fails to add new item id=%d (%v)", idx, err)
		log.Println(msg)
		return
	}

	msg = fmt.Sprintf("Item added id=%d", idx)
}
