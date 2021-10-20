package card

import (
	"errors"
)


type CardService interface {
	//Describe(card_id uint64)(string, error)
	Get(id uint64)(*Card, error)
	List(cursor uint64, limit uint64)([]Card, error)
	//Create(bank.Card)(uint64, error)
	//Update(card_id uint64, card bank.Card) error
	//Remove(card_id uint64)(bool, error)
}

type DummyCardService struct{

}

func NewDummyCardService() *DummyCardService {
	return &DummyCardService{}
}

func (p *DummyCardService) Get(id uint64)(*Card, error) {
	if int(id) >= len(allCards) {
		return nil, errors.New("out of range")
	}
	return &allCards[id], nil
}

func (p *DummyCardService) List(cursor uint64, limit uint64)([]Card, error){
	if int(cursor) >= len(allCards) {
		return nil, errors.New("out of range")
	}
	return allCards[cursor : cursor + limit], nil
}

