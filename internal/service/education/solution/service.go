package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"sort"
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
func (s *DummySolutionService) Describe(SolutionId uint64) (*education.Solution, error) {
	if _, ok := education.Data[SolutionId]; !ok {
		return nil, fmt.Errorf("Item not found")
	}
	r := education.Data[SolutionId]
	return &r, nil
}
func (s *DummySolutionService) List(cursor uint64, limit uint64) []string {
	if uint64(len(education.Data)) < cursor {
		return []string{}
	}
	if uint64(len(education.Data)) < cursor + limit {
		limit = uint64(len(education.Data)) - cursor
	}
	//Наверное есть более правильный метод, но я не смог придумать как из мапы вернуть элементы, а если через массив
	//делать то сильно усложняются другие методы
	rs := make([]string, 0, len(education.Data))
	for _, v := range education.Data {
		rs = append(rs, v.String())
	}
	sort.Strings(rs)
	res := make([]string, 0, limit)
	for i:= cursor; i < cursor + limit; i++ {
		res = append(res, rs[i])
	}
	return res
}
func (s *DummySolutionService) CreateNewID() uint64 {
	max := uint64(0)
	for i, _ := range education.Data {
		if max < i {max = i}
	}
	return max + 1
}

func (s *DummySolutionService) Create(SolutionId uint64, Solution education.Solution) (uint64, error) {
	if _, ok := education.Data[SolutionId]; ok {
		return 0, fmt.Errorf("Item already exists")
	}
	education.Data[SolutionId] = Solution
	return SolutionId, nil
}
func (s *DummySolutionService) Update(SolutionId uint64, Solution education.Solution) error {
	if _, ok := education.Data[SolutionId]; !ok {
		return fmt.Errorf("Item not found")
	}
	education.Data[SolutionId] = Solution
	return nil
}
func (s *DummySolutionService) Remove(SolutionId uint64) (bool, error) {
	if _, ok := education.Data[SolutionId]; !ok {
		return false, fmt.Errorf("Item not found")
	}
	delete(education.Data, SolutionId)
	return true, nil
}
func (s *DummySolutionService) Len() uint64 {
	return uint64(len(education.Data))
}
