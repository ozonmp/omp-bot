package schedule

import (
	"errors"
	"fmt"

	com "github.com/ozonmp/omp-bot/internal/model/travel"
)

func (s *DummyScheduleService) Update(schedule_id uint64, schedule com.Schedule) error {
	if schedule.Name == "" {
		return errors.New("name is empty")
	}

	entity, ok := s.entities[schedule_id]
	if !ok {
		return fmt.Errorf(entityNotFound, schedule_id)
	}

	entity.Name = schedule.Name
	s.entities[schedule_id] = entity
	return nil
}
