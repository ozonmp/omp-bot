package serial

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/cinema/serial"
)

func (c *CinemaSerialCommander) Edit(inputMessage *tgbotapi.Message) {
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

	idx, err := strconv.Atoi(argss[0])
	if err != nil {
		msg = fmt.Sprintf("Err: incorrect id (%v)", err)
		log.Println(msg)
		return
	}

	var new serial.Serial
	for _, v := range argss[1:] {
		if m, _ := regexp.MatchString(`^[iI][dD]=\d+$`, v); m {
			new.ID, _ = strconv.Atoi(strings.SplitN(v, "=", 2)[1])
			continue
		}
		if m, _ := regexp.MatchString(`^[tT][iI][tT][lL][eE]=\w+$`, v); m {
			new.Title = strings.SplitN(v, "=", 2)[1]
			continue
		}
		if m, _ := regexp.MatchString(`^[yY][eE][aA][rR]=\d+$`, v); m {
			new.Year, _ = strconv.Atoi(strings.SplitN(v, "=", 2)[1])
			continue
		}
	}

	if new.ID == 0 && new.Title == "" && new.Year == 0 {
		msg = "Err: no correct parameters specified. Use:\n/edit__cinema__serial <id to edit>\nid=<new id>\ntitle=<new title>\nyear=<new year>"
		log.Println(msg)
		return
	}

	err = c.subdomainService.Edit(idx, new)
	if err != nil {
		msg = fmt.Sprintf("Err: fails to edit serial id=%d (%v)", idx, err)
		log.Println(msg)
		return
	}

	msg = fmt.Sprintf("Item updated id=%d", idx)
}
