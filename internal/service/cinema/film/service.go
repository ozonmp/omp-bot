package film

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"github.com/ozonmp/omp-bot/internal/utils/logger"
	cnmApi "github.com/ozonmp/omp-bot/pb/github.com/ozonmp/cnm-film-api/pkg/cnm-film-api"
)

type DummyFilmService struct {
	filmApi cnmApi.CnmFilmApiServiceClient
}

func NewDummyFilmService(film cnmApi.CnmFilmApiServiceClient) *DummyFilmService {
	newService := &DummyFilmService{filmApi: film}
	return newService
}

func (s *DummyFilmService) Describe(ctx context.Context, filmID int64) (*cinema.Film, error) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("FilmService - Describe - Start")

	req := &cnmApi.DescribeFilmV1Request{FilmId: filmID}
	log.Debug().Msg(fmt.Sprintf("FilmService - Describe - Try DescribeFilmV1 with %v", req))
	resp, err := s.filmApi.DescribeFilmV1(ctx, req)
	if err != nil {
		log.Debug().Err(err).Msg("FilmService - Describe - DescribeFilmV1 handle error")
		return nil, err
	}

	log.Debug().Msg(fmt.Sprintf("FilmService - Describe - Got DescribeFilmV1 response %v", resp))

	pbFilm := resp.GetFilm()
	if pbFilm == nil {
		err := fmt.Errorf("empty film field in response")
		log.Debug().Err(err).Msg("FilmService - Describe - Error")
		return nil, err
	}

	film := fromPbToFilm(pbFilm)
	return &film, nil
}

func (s *DummyFilmService) List(ctx context.Context, cursor, limit int64) ([]cinema.Film, error) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("FilmService - List - Start")

	req := &cnmApi.ListFilmV1Request{
		Cursor: cursor,
		Limit: limit,
	}
	log.Debug().Msg(fmt.Sprintf("FilmService - List - Try ListFilmV1 with %v", req))
	resp, err := s.filmApi.ListFilmV1(ctx, req)
	if err != nil {
		log.Debug().Err(err).Msg("FilmService - List - ListFilmV1 handle error")
		return nil, err
	}

	log.Debug().Msg(fmt.Sprintf("FilmService - List - Got ListFilmV1 response %v", resp))

	films := fromPbToFilms(resp.GetFilm())
	return films, nil
}

func (s *DummyFilmService) Create(ctx context.Context, film *cinema.Film) (*int64, error) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("FilmService - Create - Start")

	req := &cnmApi.CreateFilmV1Request{
		Name: film.Name,
		Rating: film.Rating,
		ShortDescription: film.ShortDescription,
	}
	log.Debug().Msg(fmt.Sprintf("FilmService - Create - Try CreateFilmV1 with %v", req))
	resp, err := s.filmApi.CreateFilmV1(ctx, req)
	if err != nil {
		log.Debug().Err(err).Msg("FilmService - Create - CreateFilmV1 handle error")
		return nil, err
	}

	log.Debug().Msg(fmt.Sprintf("FilmService - Create - Got CreateFilmV1 response %v", resp))

	pbFilm := resp.GetFilm()
	if pbFilm == nil {
		err := fmt.Errorf("empty film field in response")
		log.Debug().Err(err).Msg("FilmService - Describe - Error")
		return nil, err
	}

	resFilm := fromPbToFilm(pbFilm)
	return &resFilm.ID, nil
}

func (s *DummyFilmService) Update(ctx context.Context, film *cinema.Film) (*int64, error) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("FilmService - Update - Start")

	reqFilm := fromFilmToReq(*film)
	req := &cnmApi.UpdateFilmV1Request{
		Film: &reqFilm,
	}
	log.Debug().Msg(fmt.Sprintf("FilmService - Update - Try UpdateFilmV1 with %v", req))
	resp, err := s.filmApi.UpdateFilmV1(ctx, req)
	if err != nil {
		log.Debug().Err(err).Msg("FilmService - Update - UpdateFilmV1 handle error")
		return nil, err
	}

	log.Debug().Msg(fmt.Sprintf("FilmService - Update - Got UpdateFilmV1 response %v", resp))

	pbFilm := resp.GetFilm()
	if pbFilm == nil {
		err := fmt.Errorf("empty film field in response")
		log.Debug().Err(err).Msg("FilmService - Describe - Error")
		return nil, err
	}

	resFilm := fromPbToFilm(pbFilm)
	return &resFilm.ID, nil
}

func (s *DummyFilmService) Remove(ctx context.Context, filmID int64) (*int64, error) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("FilmService - Remove - Start")

	req := &cnmApi.RemoveFilmV1Request{
		FilmId: filmID,
	}
	log.Debug().Msg(fmt.Sprintf("FilmService - Remove - Try RemoveFilmV1 with %v", req))
	resp, err := s.filmApi.RemoveFilmV1(ctx, req)
	if err != nil {
		log.Debug().Err(err).Msg("FilmService - Remove - RemoveFilmV1 handle error")
		return nil, err
	}

	log.Debug().Msg(fmt.Sprintf("FilmService - Remove - Got RemoveFilmV1 response %v", resp))

	pbFilm := resp.GetFilm()
	if pbFilm == nil {
		err := fmt.Errorf("empty film field in response")
		log.Debug().Err(err).Msg("FilmService - Remove - Error")
		return nil, err
	}

	resFilm := fromPbToFilm(pbFilm)
	return &resFilm.ID, nil
}
