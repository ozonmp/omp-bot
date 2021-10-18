package film

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"strconv"
	"strings"
)

func getParameters(text string) ([][]string, error) {
	params := make([][]string, 0)
	entityArgs := strings.Split(text, ";")
	for _, arg := range entityArgs {
		param := strings.SplitN(arg, ":", 2)
		if len(param) != 2 {
			return nil, fmt.Errorf("Bad input, use ':' as separator between field and value. Example: name:Harry potter;rating:10;description:good story")
		}
		params = append(params, []string{strings.TrimSpace(param[0]), strings.TrimSpace(param[1])})
	}
	return params, nil
}

func filmFromParameters(params [][]string) (*cinema.Film, error) {
	film := cinema.Film{Name: "", Rating: 0, ShortDescription: ""}
	var err error
	for _, param := range params {
		lowParam := strings.ToLower(param[0])
		switch lowParam {
		case "name":
			film.Name = param[1]
		case "description":
			film.ShortDescription = param[1]
		case "rating":
			film.Rating, err = strconv.ParseFloat(param[1], 64)
			if err != nil {
				return nil, fmt.Errorf("Bad rating format '%s'", param[1])
			}
		}
	}
	return &film, nil
}
