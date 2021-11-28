package film

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"strconv"
	"strings"
)

func (c CinemaFilmCommander) Edit(ctx context.Context, inputMsg *tgbotapi.Message) {
	entityArgsString := strings.TrimSpace(inputMsg.CommandArguments())
	newFilm, err := c.editCommandLogic(ctx, entityArgsString)
	if err != nil {
		_ = c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}
	_ = c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Updated film:\n%s", newFilm)))
}

func (c *CinemaFilmCommander) editCommandLogic(ctx context.Context, entityArgsString string) (*cinema.Film, error) {
	params, err := getParameters(entityArgsString)
	if err != nil {
		return nil, err
	}

	strIDToFind := ""
	for _, arg := range params {
		if arg[0] == "search" { // search:film name to find
			strIDToFind = arg[1]
			break
		}
	}

	idToFind, err := strconv.Atoi(strIDToFind)
	if err != nil || idToFind < 0 {
		return nil, fmt.Errorf("Message must contain 'search' argument with film ID, to find a film for update.\nExample: search:10;name:Harry Potter;rating:9.8;description:good story")
	}

	newFilm, err := filmFromParameters(params)
	if err != nil {
		return nil, err
	}

	newFilm.ID = int64(idToFind)
	if _, err := c.filmService.Update(ctx, newFilm); err != nil {
		return nil, err
	}

	return newFilm, nil
}
