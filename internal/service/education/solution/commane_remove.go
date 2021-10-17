package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

func (s *DummySolutionService) Remove(SolutionId uint64) (bool, error) {
	if _, ok := education.Data[SolutionId]; !ok {
		return false, fmt.Errorf("item not found")
	}
	delete(education.Data, SolutionId)
	return true, nil
}

