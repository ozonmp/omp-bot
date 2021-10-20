package provider

import (
	"fmt"
	"log"
)

type ProviderService interface {
	Get(providerID uint64) (*Provider, error)
	Create(Provider) (uint64, error)
	Remove(providerID uint64) (bool, error)
	Update(providerID uint64, provider Provider) error
	List(cursor uint64, limit uint64) []Provider
	EntitiesCount() int
	ShortDescription(provider *Provider) string
	LongDescription(provider *Provider) string
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}
func (s *Service) EntitiesCount() int { return len(allEntities) }
func (s *Service) List(cursor uint64, limit uint64) []Provider {
	if limit == 0 {
		return allEntities
	}
	lenI64 := uint64(len(allEntities))
	if cursor+limit >= lenI64 {
		return allEntities[cursor:]
	}
	return allEntities[cursor : limit+cursor]
}
func (s *Service) Create(provider Provider) (uint64, error) {
	lastEnt := allEntities[len(allEntities)-1] //supposed to have maximum id value for current implementation
	newId := lastEnt.Id + 1
	provider.Id = newId
	allEntities = append(allEntities, provider)
	return newId, nil
}
func (s *Service) Remove(providerID uint64) (bool, error) {
	for i := 0; i < len(allEntities); i++ {
		if providerID == allEntities[i].Id {
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}
func (s *Service) Update(providerID uint64, provider Provider) error {
	p, err := s.Get(providerID)
	if err != nil {
		return err
	}
	p.Name = provider.Name
	p.Description = provider.Description
	p.IsImplemented = provider.IsImplemented

	return nil
}
func (s *Service) Get(providerID uint64) (*Provider, error) {
	for i := 0; i < len(allEntities); i++ {
		if providerID == allEntities[i].Id {
			return &allEntities[uint64(i)], nil
		}
	}
	err := fmt.Errorf("provider with ID %d not found", providerID)
	log.Printf("provider.Service.Get: %v", err)
	return nil, err
}
func (s *Service) ShortDescription(provider *Provider) string {
	return fmt.Sprintf("%d | %s",
		provider.Id,
		provider.Name)
}
func (s *Service) LongDescription(provider *Provider) string {
	return fmt.Sprintf(
		"ID : %d\nName: %s\nDescription: %s\nIsImplemented: %v",
		provider.Id,
		provider.Name,
		provider.Description,
		provider.IsImplemented)
}
