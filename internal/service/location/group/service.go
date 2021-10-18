package group

import "github.com/ozonmp/omp-bot/internal/model/location"

//type GroupService interface {
//	List(cursor uint64, limit uint64) ([]location.Group, error)
//}

type LocationGroupService struct {
	allEntities []location.Group
}

func NewLocationGroupService() *LocationGroupService {
	return &LocationGroupService{
		allEntities: []location.Group{
			{Id: 1, CountOfLocation: 2, Type: "building"},
			{Id: 2, CountOfLocation: 3, Type: "building"},
			{Id: 3, CountOfLocation: 3, Type: "room"},
			{Id: 4, CountOfLocation: 1, Type: "warehouse"},
			{Id: 5, CountOfLocation: 1, Type: "warehouse"},
		},
	}
}

func (s *LocationGroupService) List() []location.Group {
	return s.allEntities
}
