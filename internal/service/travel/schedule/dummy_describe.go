package schedule

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/travel"
)

func (s *DummyScheduleService) Describe(schedule_id uint64) (*travel.Schedule, error) {
	entitiy, ok := s.entities[schedule_id]
	if !ok {
		return nil, fmt.Errorf(entityNotFound, schedule_id)
	}

	return &entitiy, nil
}
