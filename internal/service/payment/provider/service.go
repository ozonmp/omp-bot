package provider

import (
	"fmt"
	"log"
)

type ProviderService interface {
	Describe(providerID uint64) (*Provider, error)
	Create(Provider) (uint64, error)
	Remove(providerID uint64) (bool, error)
	Update(providerID uint64, provider Provider) error
	List(cursor uint64, limit uint64) []Provider
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(cursor uint64, limit uint64) []Provider {
	return allEntities
}
func (s *Service) Create(provider Provider) (uint64, error) {
	log.Println(provider)
	return 0, nil
}
func (s *Service) Remove(providerID uint64) (bool, error) {
	return false, nil
}
func (s *Service) Update(providerID uint64, provider Provider) error {
	return nil
}

func (s *Service) Describe(providerID uint64) (*Provider, error) {
	for i := 0; i < len(allEntities); i++ {
		if providerID == allEntities[i].Id {
			return &allEntities[uint64(i)], nil
		}
	}
	err := fmt.Errorf("provider with ID %d not found", providerID)
	log.Printf("provider.Service.Describe: %v", err)
	return nil, err
}
