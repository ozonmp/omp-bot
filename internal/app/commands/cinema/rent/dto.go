package rent

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

type InputItemDTO struct {
	FilmID   int64 `json:"film_id"`
	SerialID int64 `json:"serial_id"`
	Price    int64 `json:"price"`
}

type ListLimitDTO struct {
	Cursor int64 `json:"cursor"`
	Limit  int64 `json:"limit"`
}

func (c *CinemaRentCommander) jsonInputParser(inputMessage *tgbotapi.Message) (*cinema.Rent, error) {
	args := strings.Trim(inputMessage.CommandArguments(), " ")

	dto := &InputItemDTO{FilmID: -1, SerialID: -1, Price: -1}

	if err := json.Unmarshal([]byte(args), dto); err != nil {
		log.Printf("CinemaRentCommander.jsonInputParser: %v", err)
		return nil, err
	}

	if dto.Price < 0 {
		return nil, fmt.Errorf("CinemaRentCommander.jsonInputParser: Не передано корректное значение стоимости")
	}
	if dto.FilmID == -1 && dto.SerialID == -1 {
		return nil, fmt.Errorf("CinemaRentCommander.jsonInputParser: Не передано корректное значение идентификаторов фильма/сериала")
	}

	rent := &cinema.Rent{
		FilmID:       dto.FilmID,
		SerialID:     dto.SerialID,
		PriceInKopec: dto.Price,
	}

	return rent, nil
}

func (cinema *CinemaRentCommander) jsonListCallbackParser(callbackPath path.CallbackPath) (*ListLimitDTO, error) {
	data := &ListLimitDTO{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
