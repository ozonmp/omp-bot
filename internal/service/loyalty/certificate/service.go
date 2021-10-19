package certificate

import "errors"

type СertificateService interface {
	Describe(certificateID uint64) (*Сertificate, error)
	List(cursor uint64, limit uint64) ([]Сertificate, error)
	Create(certificate Сertificate) (uint64, error)
	Update(certificateID uint64, certificate Сertificate) error
	Remove(certificateID uint64) (bool, error)
}

type DummyСertificateService struct {

}

func NewDummyСertificateService() *DummyСertificateService {
	return &DummyСertificateService{}
}

func (s *DummyСertificateService) Describe(certificateID uint64) (*Сertificate, error) {
	return nil, nil
}

func (s *DummyСertificateService) List(cursor uint64, limit uint64) ([]Сertificate, error) {
	size := uint64(len(allEntities))

	if cursor >= size {
		return nil, errors.New("cursor is greater then array size")
	}
	if cursor + limit >= size {
		limit = size - cursor
	}

	return allEntities[cursor:cursor+limit], nil
}

func (s *DummyСertificateService) Get(idx int) (*Сertificate, error) {
	if idx >= len(allEntities) {
		return nil, errors.New("index is out of bounds")
	}
	return &allEntities[idx], nil
}
