package certificate

import (
	"errors"
	"fmt"
)

type СertificateService interface {
	Describe(certificateID uint64) (*Certificate, error)
	List(cursor uint64, limit uint64) ([]Certificate, error)
	Create(certificate Certificate) (uint64, error)
	Update(certificateID uint64, certificate Certificate) error
	Remove(certificateID uint64) (bool, error)
}

type DummyСertificateService struct {
	Certificates []Certificate
}

func NewDummyСertificateService() *DummyСertificateService {
	return &DummyСertificateService{ Certificates: allEntities }
}

func (s *DummyСertificateService) Describe(certificateID uint64) (*Certificate, error) {
	for _, certificate := range s.Certificates {
		if certificate.Id == certificateID {
			return &certificate, nil
		}
	}
	return nil, errors.New("id not found")
}

func (s *DummyСertificateService) List(cursor uint64, limit uint64) ([]Certificate, error) {
	size := uint64(len(s.Certificates))
	from := cursor
	to := cursor + limit

	if from > size {
		from = size
	}
	if to > size {
		to = size
	}

	return s.Certificates[cursor:to], nil
}

func (s *DummyСertificateService) Remove(certificateID uint64) (bool, error) {
	for i, certificate := range s.Certificates {
		if certificate.Id == certificateID {
			s.Certificates = append(s.Certificates[:i], s.Certificates[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (s *DummyСertificateService) Create(newCertificate Certificate) (uint64, error) {

	for _, certificate := range s.Certificates {
		if certificate.Id == newCertificate.Id {
			return newCertificate.Id, errors.New(fmt.Sprintf("certificate with ID %d already exists", newCertificate.Id))
		}
	}
	s.Certificates = append(s.Certificates[:], newCertificate)
	return newCertificate.Id, nil
}