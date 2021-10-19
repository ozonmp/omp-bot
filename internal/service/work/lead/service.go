package lead

import (
	"fmt"
	"sort"
)

type DummyLeadService struct{}

func NewDummyLeadService() *DummyLeadService {
	return &DummyLeadService{}
}

func (s *DummyLeadService) List(offset uint64, limit uint64) ([]Lead, error) {
	var totalLeads = uint64(len(allEntities))
	if offset >= totalLeads {
		return []Lead{}, nil
	}
	res := make([]Lead, 0, totalLeads)
	for _, lead := range allEntities {
		res = append(res, lead)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})

	end := offset + limit
	if end > totalLeads {
		end = totalLeads
	}

	return res[offset:end], nil
}

func (s *DummyLeadService) Describe(leadID uint64) (*Lead, error) {
	lead, exists := allEntities[leadID]
	if !exists {
		return nil, fmt.Errorf(`Lead #%d not found`, leadID)
	}
	return &lead, nil
}

func (s *DummyLeadService) Create(lead Lead) (uint64, error) {
	if lead.FirstName == "" || lead.LastName == "" || lead.Project == "" {
		return 0, fmt.Errorf(`Validation error`)
	}

	nextID := s.getNextId()
	lead.ID = nextID
	allEntities[nextID] = lead

	return nextID, nil
}

func (s *DummyLeadService) Update(leadID uint64, lead Lead) error {
	exLead, exists := allEntities[leadID]
	if !exists {
		return fmt.Errorf(`Lead #%d not found`, leadID)
	}

	if lead.FirstName == "" || lead.LastName == "" || lead.Project == "" {
		return fmt.Errorf(`Validation error`)
	}

	exLead.FirstName = lead.FirstName
	exLead.LastName = lead.LastName
	exLead.Project = lead.Project

	allEntities[leadID] = exLead

	return nil
}

func (s *DummyLeadService) Remove(leadID uint64) (bool, error) {
	_, exists := allEntities[leadID]
	if !exists {
		return false, fmt.Errorf(`Lead #%d not found`, leadID)
	}
	delete(allEntities, leadID)
	return true, nil
}

func (s *DummyLeadService) getNextId() uint64 {
	idSeq++
	return idSeq
}
