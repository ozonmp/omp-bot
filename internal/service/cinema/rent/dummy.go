package rent

import (
	"fmt"
	"time"

	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

type DummyRentService struct {
	items []cinema.Rent
}

func NewDummyRentService() *DummyRentService {
	service := &DummyRentService{
		items: make([]cinema.Rent, 0),
	}

	service.Create(*cinema.NewFilmRent(1, 1000))
	service.Create(*cinema.NewFilmRent(2, 70))
	service.Create(*cinema.NewFilmRent(3, 32556))
	service.Create(*cinema.NewFilmRent(4, 8345))
	service.Create(*cinema.NewFilmRent(5, 1000))
	service.Create(*cinema.NewFilmRent(6, 70))
	service.Create(*cinema.NewFilmRent(7, 32556))
	service.Create(*cinema.NewFilmRent(8, 8345))
	service.Create(*cinema.NewFilmRent(9, 32556))
	service.Create(*cinema.NewFilmRent(10, 8345))

	service.Create(*cinema.NewSerialRent(1, 54656))
	service.Create(*cinema.NewSerialRent(2, 4554))
	service.Create(*cinema.NewSerialRent(3, 1245))
	service.Create(*cinema.NewSerialRent(4, 5624))
	service.Create(*cinema.NewSerialRent(5, 54656))
	service.Create(*cinema.NewSerialRent(6, 4554))
	service.Create(*cinema.NewSerialRent(7, 1245))
	service.Create(*cinema.NewSerialRent(8, 5624))
	service.Create(*cinema.NewSerialRent(9, 1245))
	service.Create(*cinema.NewSerialRent(10, 5624))

	return service
}

func (s *DummyRentService) Describe(rentID uint64) (*cinema.Rent, error) {
	if rentID > uint64(len(s.items)-1) {
		return nil, fmt.Errorf("DummyRentService.Describe: запись с идентификатором %d не найдена", rentID)
	}

	if s.items[rentID].Deleted == true {
		return nil, fmt.Errorf(
			"DummyRentService.Describe: запись с идентифкатором %d была удалена в %s",
			rentID,
			s.items[rentID].DeletedAt.Format(time.RFC1123),
		)
	}

	return &s.items[rentID], nil
}

func (s *DummyRentService) List(cursor uint64, limit uint64) ([]cinema.Rent, error) {
	response := []cinema.Rent{}

	var index uint64 = 0
	for _, item := range s.items {
		// если элемент удалён, пропускаем такт
		if item.Deleted {
			continue
		}
		// инкрементируем реальный индекс
		index += 1

		// если мы ещё не достигли элемента, с которого должны делать выборку - пропускаем
		if index < cursor {
			continue
		}
		if index > cursor+limit {
			break
		}

		response = append(response, item)
	}

	return response, nil
}

func (s *DummyRentService) Create(item cinema.Rent) (uint64, error) {
	if item.FilmID != -1 && item.SerialID != -1 {
		return 0, fmt.Errorf("DummyRentService.Create: Добавляемый элемент не может содержать одновременно идентификаторы фильма и сериала")
	}

	if item.FilmID != -1 {
		return s.pushFilm(item)
	} else if item.SerialID != -1 {
		return s.pushSerial(item)
	} else {
		return 0, fmt.Errorf("DummyRentService.Create: Не передан ни один из идентификаторов фильма или сериала")
	}
}

func (s *DummyRentService) pushFilm(item cinema.Rent) (uint64, error) {
	if s.FilmIDExists(item.FilmID) {
		return 0, fmt.Errorf("DummyRentService.pushFilm: Добавляемый фильм с идентифкатором %d уже существует", item.FilmID)
	}

	// UglyHack. созраним индекс под которым будет содержаться запись
	item.RecordIndex = uint64(len(s.items))

	s.items = append(s.items, item)
	return item.RecordIndex, nil
}

func (s *DummyRentService) pushSerial(item cinema.Rent) (uint64, error) {
	if s.SerialIDExists(item.SerialID) {
		return 0, fmt.Errorf("DummyRentService.pushSerial: Добавляемый сериал с идентифкатором %d уже существует", item.SerialID)
	}

	// UglyHack. созраним индекс под которым будет содержаться запись
	item.RecordIndex = uint64(len(s.items))

	s.items = append(s.items, item)
	return item.RecordIndex, nil
}

func (s *DummyRentService) Update(rentID uint64, rent cinema.Rent) error {
	if rentID > uint64(len(s.items)) {
		return fmt.Errorf("DummyRentService.Update: Запись с идентификатором %d не найдена", rentID)
	}

	if s.items[rentID].Deleted == true {
		return fmt.Errorf(
			"DummyRentService.Update: Запись с идентификатором %d была удалена в %s",
			rentID,
			s.items[rentID].DeletedAt.Format(time.RFC1123),
		)
	}

	if rent.FilmID != -1 && rent.SerialID != -1 {
		return fmt.Errorf("DummyRentService.Update: Запись не может содержать одновременно идентификаторы фильма и сериала")
	}

	if rent.FilmID != -1 {
		return s.updateFilm(rentID, rent)
	} else if rent.SerialID != -1 {
		return s.updateSerial(rentID, rent)
	} else {
		return fmt.Errorf("DummyRentService.Update: Не передан ни один из идентификаторов фильма или сериала")
	}
}

func (s *DummyRentService) updateFilm(rentID uint64, rent cinema.Rent) error {
	if s.items[rentID].FilmID != rent.FilmID && s.FilmIDExists(rent.FilmID) {
		return fmt.Errorf("DummyRentService.updateFilm: уже существует другая запись с идентификатором фильма %d", rent.FilmID)
	}

	s.items[rentID] = rent

	return nil
}

func (s *DummyRentService) updateSerial(rentID uint64, rent cinema.Rent) error {
	if s.items[rentID].SerialID != rent.SerialID && s.SerialIDExists(rent.SerialID) {
		return fmt.Errorf("DummyRentService.updateSerial: уже существует другая запись с идентификатором сериала %d", rent.SerialID)
	}

	s.items[rentID] = rent

	return nil
}

func (s *DummyRentService) Remove(rentID uint64) (bool, error) {
	item, err := s.Describe(rentID)
	if err != nil {
		return false, err
	}

	item.Deleted = true
	item.DeletedAt = time.Now()

	return true, nil
}

func (s *DummyRentService) getItemIndexWithFilmID(filmID int64) (uint64, error) {
	for index, item := range s.items {
		if item.FilmID == filmID && item.Deleted == false {
			return uint64(index), nil
		}
	}
	return 0, fmt.Errorf("Элемент с идентификатором фильма %d не найден", filmID)
}

func (s *DummyRentService) getItemIndexWithSerialID(SerialID int64) (uint64, error) {
	for index, item := range s.items {
		if item.SerialID == SerialID && item.Deleted == false {
			return uint64(index), nil
		}
	}
	return 0, fmt.Errorf("Элемент с идентификатором сериала %d не найден", SerialID)
}

func (s *DummyRentService) FilmIDExists(FilmID int64) bool {
	for _, item := range s.items {
		if item.FilmID == FilmID && item.Deleted == false {
			return true
		}
	}
	return false
}

func (s *DummyRentService) SerialIDExists(SerialID int64) bool {
	for _, item := range s.items {
		if item.SerialID == SerialID && item.Deleted == false {
			return true
		}
	}
	return false
}
