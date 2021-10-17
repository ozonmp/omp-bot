package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

func (s *DummySolutionService) Describe(SolutionId uint64) (*education.Solution, error) {
	if _, ok := education.Data[SolutionId]; !ok {
		return nil, fmt.Errorf("item not found")
	}
	r := education.Data[SolutionId]
	return &r, nil
}

