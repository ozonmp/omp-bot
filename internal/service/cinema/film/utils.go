package film

import (
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	filmApi "github.com/ozonmp/omp-bot/pb/github.com/ozonmp/cnm-film-api/pkg/cnm-film-api"
)

func fromFilmToReq(film cinema.Film) filmApi.ReqFilm {
	return filmApi.ReqFilm{
		Id: film.ID,
		Name: film.Name,
		Rating: film.Rating,
		ShortDescription: film.ShortDescription,
	}
}

func fromPbToFilm(pbFilm *filmApi.Film) cinema.Film {
	return cinema.Film{
		ID: pbFilm.Id,
		Name: pbFilm.Name,
		Rating: pbFilm.Rating,
		ShortDescription: pbFilm.ShortDescription,
	}
}

func fromPbToFilms(pbFilms []*filmApi.Film) []cinema.Film {
	films := make([]cinema.Film, len(pbFilms))
	for idx, pbFilm := range pbFilms {
		films[idx] = fromPbToFilm(pbFilm)
	}
	return films
}
