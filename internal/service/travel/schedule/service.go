package schedule

import (
	"github.com/ozonmp/omp-bot/internal/model/travel"
)

const entityNotFound = "entity %d not found"

type ScheduleService interface {
	Describe(schedule_id uint64) (*travel.Schedule, error)
	List(cursor uint64, limit uint64) ([]travel.Schedule, error)
	Create(travel.Schedule) (uint64, error)
	Update(schedule_id uint64, kjubybot travel.Schedule) error
	Remove(schedule_id uint64) (bool, error)
}

type DummyScheduleService struct {
	entities map[uint64]travel.Schedule
	lastID   uint64
}

func NewDummyScheduleService() *DummyScheduleService {
	return &DummyScheduleService{
		make(map[uint64]travel.Schedule),
		0,
	}
}
