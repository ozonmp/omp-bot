package schedule

import "fmt"

func (s *DummyScheduleService) Remove(schedule_id uint64) (bool, error) {
	_, ok := s.entities[schedule_id]
	if !ok {
		return false, fmt.Errorf(entityNotFound, schedule_id)
	}

	delete(s.entities, schedule_id)
	if schedule_id == s.lastID {
		s.lastID--
	}

	return true, nil
}
