package transition

import (
	"fmt"
	"sync/atomic"

	"github.com/ozonmp/omp-bot/internal/model/activity"
)

type TransitionService interface {
	Describe(transitionID uint64) (*activity.Transition, error)
	List(offset uint64, limit uint64) ([]activity.Transition, error)
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

func (s *DummyTransitionService) List(offset uint64, limit uint64) ([]activity.Transition, error) {
	transitions := make([]activity.Transition, 0, limit)
	var index uint64 = 0
	var counter uint64 = 0

	// skip elements before offset
	for index = 0; index <= s.lastIndex; index++ {
		if counter == offset {
			break
		}
		_, ok := allEntities[index]
		if !ok {
			continue
		} else {
			counter++
		}
	}

	// get elements from offset to offset + limit
	for ; index <= s.lastIndex; index++ {
		if len(transitions) == int(limit) {
			break
		}

		transition, ok := allEntities[index]
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
