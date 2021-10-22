package referral

import (
	"errors"
	"log"
)

type ReferralService interface {
	Describe(referralID uint64) (*Referral, error)
	List(cursor uint64, limit uint64) ([]Referral, error)
	Create(Referral) (uint64, error)
	Update(referralID uint64, referral Referral) error
	Remove(referralID uint64) (bool, error)
}

type DummyReferralService struct{}

func NewDummyDummyReferralService() *DummyReferralService {
	return &DummyReferralService{}
}

func (s *DummyReferralService) Describe(referralId uint64) (*Referral, error) {
	for i := 0; i < len(tempReferrals); i++ {
		if referralId == tempReferrals[i].id {
			return &tempReferrals[i], nil
		}
	}
	err := errors.New("Invalid referralId")
	return nil, err
}

func (s *DummyReferralService) List(cursor uint64, limit uint64) ([]Referral, error) {
	referralLength := uint64(len(tempReferrals))
	if cursor >= referralLength {
		return nil, errors.New("invalid cursor")
	}

	if cursor+limit < referralLength {
		return tempReferrals[cursor : limit+cursor], nil
	} else {
		return tempReferrals[cursor:], nil
	}

}

func (s *DummyReferralService) Create(referral Referral) (uint64, error) {
	var max uint64 = 0
	for i := 0; i < len(tempReferrals); i++ {
		if max < tempReferrals[i].id {
			max = tempReferrals[i].id
		}
	}
	referral.id = max + 1
	tempReferrals = append(tempReferrals, referral)
	log.Print(tempReferrals)
	return referral.id, nil
}

func (s *DummyReferralService) Remove(referralId uint64) (bool, error) {
	for i := 0; i < len(tempReferrals); i++ {
		if referralId == tempReferrals[i].id {
			tempReferrals = append(tempReferrals[:i], tempReferrals[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (s *DummyReferralService) Update(referral Referral) error {
	for i := 0; i < len(tempReferrals); i++ {
		if referral.id == tempReferrals[i].id {
			tempReferrals[i] = referral
			return nil
		}
	}
	return errors.New("invalid id")
}
