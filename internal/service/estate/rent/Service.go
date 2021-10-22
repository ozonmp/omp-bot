package rent

import (
	"github.com/ozonmp/omp-bot/internal/model/estate"
	"github.com/ozonmp/omp-bot/internal/service/estate/rent/storage"
)

type Service struct {
	storage *storage.MemoryStorage
}

func NewService(storage *storage.MemoryStorage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Describe(rentID uint64) (*estate.Rent, error) {
	if entity, err := s.storage.Read(rentID); err == nil {
		model := entity.ToModel()
		return &model, nil
	} else {
		return nil, err
	}
}

func (s *Service) List(cursor uint64, limit uint64) ([]estate.Rent, error) {
	if cursor == 0 && limit == 0 {
		return translate(s.storage.ReadAll()), nil
	}

	count := s.storage.Count()

	if count == 0 {
		return translate(s.storage.ReadAll()), nil
	}

	last := cursor + limit
	if last > count {
		limit -= last - count
	}

	if arr, err := s.storage.ReadPage(cursor, limit); err == nil {
		return translate(arr), nil
	} else {
		return nil, err
	}
}

func (s *Service) Create(model estate.Rent) (uint64, error) {
	return s.storage.Create(storage.NewRentEntity(0, model))
}

func (s *Service) Update(rentID uint64, rent estate.Rent) error {
	entity := storage.NewRentEntity(rentID, rent)
	return s.storage.Update(entity)
}

func (s *Service) Remove(rentID uint64) (bool, error) {
	return s.storage.Delete(rentID)
}

func translate(arr []*storage.RentEntity) []estate.Rent {
	a := make([]estate.Rent, 0, len(arr))
	for _, v := range arr {
		a = append(a, v.ToModel())
	}
	return a
}
