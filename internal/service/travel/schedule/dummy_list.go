package schedule

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/travel"
)

func (s *DummyScheduleService) List(cursor uint64, limit uint64) ([]travel.Schedule, error) {
	if len(s.entities) == 0 {
		return []travel.Schedule{}, nil
	}

	if cursor > s.lastID {
		return nil, fmt.Errorf(entityNotFound, cursor)
	}

	result := make([]travel.Schedule, 0, limit)
	for i := cursor; i <= s.lastID && limit > 0; i++ {
		if k, ok := s.entities[i]; ok {
			result = append(result, k)
			limit--
		}
	}

	return result, nil
}
