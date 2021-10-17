package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

func (s *DummySolutionService) Update(SolutionId uint64, Solution education.Solution) error {
	if _, ok := education.Data[SolutionId]; !ok {
		return fmt.Errorf("item not found")
	}
	education.Data[SolutionId] = Solution
	return nil
}

