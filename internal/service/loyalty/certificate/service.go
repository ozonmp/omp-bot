package certificate

import "errors"

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
	size := uint64(len(allEntities))
	from := cursor
	to := cursor + limit

	if from > size {
		from = size
	}
	if to > size {
		to = size
	}

	return allEntities[cursor:to], nil
}
