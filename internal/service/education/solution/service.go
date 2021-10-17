package solution

import (
	"github.com/ozonmp/omp-bot/internal/model/education"
)

type SolutionService interface {
	Describe(SolutionId uint64) (*education.Solution, error)
	List(cursor uint64, limit uint64) []string
	Create(education.Solution) (uint64, error)
	Update(SolutionId uint64, Solution education.Solution) error
	Remove(SolutionId uint64) (bool, error)
}

type DummySolutionService struct {
}

func NewDummySolutionService() *DummySolutionService {
	return &DummySolutionService{}
}

