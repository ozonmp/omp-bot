package schedule

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/travel"
)

func (s *DummyScheduleService) Create(schedule travel.Schedule) (uint64, error) {
	if schedule.Name == "" {
		return 0, errors.New("name is empty")
	}
	s.lastID++
	schedule.ID = s.lastID
	s.entities[s.lastID] = schedule
	return s.lastID, nil
}
