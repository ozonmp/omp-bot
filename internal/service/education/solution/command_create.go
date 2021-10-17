package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

func (s *DummySolutionService) Create(SolutionId uint64, Solution education.Solution) (uint64, error) {
	if _, ok := education.Data[SolutionId]; ok {
		return 0, fmt.Errorf("item already exists")
	}
	education.Data[SolutionId] = Solution
	return SolutionId, nil
}

