package film

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

var iD uint64 = 1

var examples = []cinema.Film{{Name: "test1", Rating: 0.5, ShortDescription: "descr 1"},
	{Name: "test2", Rating: 1.2, ShortDescription: "descr 2"},
	{Name: "test3", Rating: 2.3, ShortDescription: "descr 3"},
	{Name: "test4", Rating: 3.4, ShortDescription: "descr 4"},
	{Name: "test5", Rating: 4.6, ShortDescription: "descr 5"},
	{Name: "test6", Rating: 5.9, ShortDescription: "descr 6"},
	{Name: "test7", Rating: 6.5, ShortDescription: "descr 7"}}

func (s *DummyFilmService) fillByExamples() {
	for i := range examples {
		if _, err := s.Create(&examples[i]); err != nil {
			continue
		}
	}
}

type DummyFilmService struct {
	Films []cinema.Film
}

func NewDummyFilmService() *DummyFilmService {
	newService := &DummyFilmService{}
	newService.fillByExamples()
	return newService
}

func (s *DummyFilmService) Describe(filmID uint64) (*cinema.Film, error) {
	foundFilm, err := s.findByID(filmID)
	if err != nil {
		return nil, err
	}
	return foundFilm, nil
}

func (s *DummyFilmService) List(cursor, limit uint64) ([]cinema.Film, error) {
	startIndex := cursor
	endIndex := cursor + limit
	if startIndex < 0 {
		startIndex = 0
	}

	if endIndex > uint64(len(s.Films)) {
		endIndex = uint64(len(s.Films))
	}

	return s.Films[startIndex:endIndex], nil
}

func (s *DummyFilmService) Create(film *cinema.Film) (uint64, error) {
	if err := s.checkFilm(*film); err != nil {
		return 0, err
	}

	if film.ShortDescription == "" {
		film.ShortDescription = "No description provided"
	}

	film.ID = iD
	s.Films = append(s.Films, *film)
	iD += 1
	return iD, nil
}

func (s *DummyFilmService) Update(filmID uint64, film *cinema.Film) error {
	if err := s.checkFilm(*film); err != nil {
		return err
	}

	foundFilm, err := s.findByID(filmID)
	if err != nil {
		return err
	}
	film.ID = foundFilm.ID
	*foundFilm = *film
	return nil
}

func (s *DummyFilmService) Remove(filmID uint64) (bool, error) {
	for i := range s.Films {
		if s.Films[i].ID == filmID {
			s.Films = append(s.Films[:i], s.Films[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("Cant find film with ID %d", filmID)
}

func (s *DummyFilmService) NumberOfFilms() int {
	return len(s.Films)
}

func (s *DummyFilmService) checkFilm(film cinema.Film) error {
	if film.Name == "" {
		return fmt.Errorf("Name can't be blank")
	}

	if film.Rating > 10 && film.Rating < 0 {
		return fmt.Errorf("Wrong rating value, it must be between 0 and 10")
	}

	return nil
}

func (s *DummyFilmService) findByID(filmID uint64) (*cinema.Film, error) {
	for i := range s.Films {
		if s.Films[i].ID == filmID {
			return &s.Films[i], nil
		}
	}
	return nil, fmt.Errorf("Can't find film with ID %d", filmID)
}
