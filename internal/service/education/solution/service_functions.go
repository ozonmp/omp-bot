package solution

import "github.com/ozonmp/omp-bot/internal/model/education"

func (s *DummySolutionService) CreateNewID() uint64 {
	//not threadsafe
	max := uint64(0)
	for i, _ := range education.Data {
		if max < i {max = i}
	}
	return max + 1
}

func (s *DummySolutionService) Len() uint64 {
	return uint64(len(education.Data))
}
