package author

import (
	"errors"
	"log"
)

type AuthorService interface {
	Describe(referralID uint64) (*Author, error)
	List(cursor uint64, limit uint64) ([]Author, error)
	Create(Author) (uint64, error)
	Update(authorID uint64, author Author) error
	Remove(authorID uint64) (bool, error)
}

type DummyAuthorService struct{}

func NewDummyAuthorService() *DummyAuthorService {
	return &DummyAuthorService{}
}

func (s *DummyAuthorService) Describe(authorID uint64) (*Author, error) {
	for i := 0; i < len(tempAuthors); i++ {
		if authorID == tempAuthors[i].id {
			return &tempAuthors[i], nil
		}
	}
	err := errors.New("Invalid authorID")
	return nil, err
}

func (s *DummyAuthorService) List(cursor uint64, limit uint64) ([]Author, error) {
	authorLength := uint64(len(tempAuthors))
	if cursor >= authorLength {
		return nil, errors.New("invalid cursor")
	}

	if cursor+limit < authorLength {
		return tempAuthors[cursor : limit+cursor], nil
	} else {
		return tempAuthors[cursor:], nil
	}

}

func (s *DummyAuthorService) Create(author Author) (uint64, error) {
	var max uint64 = 0
	for i := 0; i < len(tempAuthors); i++ {
		if max < tempAuthors[i].id {
			max = tempAuthors[i].id
		}
	}
	author.id = max + 1
	tempAuthors = append(tempAuthors, author)
	log.Print(tempAuthors)
	return author.id, nil
}

func (s *DummyAuthorService) Remove(authorId uint64) (bool, error) {
	for i := 0; i < len(tempAuthors); i++ {
		if authorId == tempAuthors[i].id {
			tempAuthors = append(tempAuthors[:i], tempAuthors[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (s *DummyAuthorService) Update(author Author) error {
	for i := 0; i < len(tempAuthors); i++ {
		if author.id == tempAuthors[i].id {
			tempAuthors[i] = author
			return nil
		}
	}
	return errors.New("invalid id")
}
