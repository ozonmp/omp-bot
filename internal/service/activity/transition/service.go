package transition

import (
	"fmt"
	"sync/atomic"

	"github.com/ozonmp/omp-bot/internal/model/activity"
)

type TransitionService interface {
	Describe(transitionID uint64) (*activity.Transition, error)
	List(cursor uint64, limit uint64) ([]activity.Transition, error)
	Create(activity.Transition) (uint64, error)
	Update(transitionID uint64, transition activity.Transition) error
	Remove(transitionID uint64) (bool, error)
	Size() uint64
}

type DummyTransitionService struct {
	lastIndex uint64
}

func NewDummyTransitionService() *DummyTransitionService {
	var maxNumber uint64 = 0
	for n := range allEntities {
		if n > maxNumber {
			maxNumber = n
		}
	}

	return &DummyTransitionService{
		lastIndex: maxNumber,
	}
}

func (s *DummyTransitionService) Describe(transitionID uint64) (*activity.Transition, error) {
	transition, ok := allEntities[transitionID]
	if !ok {
		return nil, fmt.Errorf("Transition with id: %d not found", transitionID)
	}
	return &transition, nil
}

func (s *DummyTransitionService) List(cursor uint64, limit uint64) ([]activity.Transition, error) {
	transitions := make([]activity.Transition, 0, limit)
	for i := cursor; i <= s.lastIndex; i++ {
		if len(transitions) >= int(limit) {
			break
		}

		transition, ok := allEntities[i]
		if !ok {
			continue
		}
		transitions = append(transitions, transition)
	}
	return transitions, nil
}

func (s *DummyTransitionService) Create(transition activity.Transition) (uint64, error) {
	id := atomic.AddUint64(&s.lastIndex, 1)
	transition.Id = id
	allEntities[id] = transition

	return id, nil
}

func (s *DummyTransitionService) Update(transitionID uint64, transition activity.Transition) error {
	transition.Id = transitionID
	_, ok := allEntities[transitionID]
	if !ok {
		return fmt.Errorf("Transition with id: %d not found", transitionID)
	}
	allEntities[transitionID] = transition
	return nil
}

func (s *DummyTransitionService) Remove(transitionID uint64) (bool, error) {
	_, ok := allEntities[transitionID]
	if !ok {
		return false, fmt.Errorf("Transition with id: %d not found", transitionID)
	}
	delete(allEntities, transitionID)
	return true, nil
}

func (s *DummyTransitionService) Size() uint64 {
	return uint64(len(allEntities))
}
