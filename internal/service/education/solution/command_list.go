package solution

import (
	"github.com/ozonmp/omp-bot/internal/model/education"
	"sort"
)

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

