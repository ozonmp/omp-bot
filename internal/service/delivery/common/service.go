package common

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/delivery"
)

type CommonService interface {
	Describe(commonId uint64) (*delivery.Common, error)
	List(cursor uint64, limit uint64) ([]delivery.Common, error)
	Create(delivery.Common) (uint64, error)
	Update(commonId uint64, common delivery.Common) error
	Remove(commonId uint64) (bool, error)
}

type DummyCommonService struct{}

func NewDummyCommonService() *DummyCommonService {
	return &DummyCommonService{}
}

func (s DummyCommonService) Describe(commonId uint64) (*delivery.Common, error) {
	index, err := storageIndexById(commonId)
	if err != nil {
		errUnknownCommonIdentifier := errors.New(fmt.Sprintf("Wrong delivery id %d", commonId))
		return &delivery.Common{}, errUnknownCommonIdentifier
	}

	return &delivery.CommonStorage[index], nil
}

func (s DummyCommonService) List(cursor uint64, limit uint64) ([]delivery.Common, error) {
	if int(cursor) > len(delivery.CommonStorage)-1 {
		return []delivery.Common{}, errors.New("cursor out of range")
	}

	high := cursor + limit
	storageLen := uint64(len(delivery.CommonStorage))
	if high > storageLen {
		high = storageLen
	}

	return delivery.CommonStorage[cursor:high], nil
}

func (s DummyCommonService) Create(common delivery.Common) (uint64, error) {
	delivery.Sequence++
	common.SetId(delivery.Sequence)
	delivery.CommonStorage = append(delivery.CommonStorage, common)

	return common.Id(), nil
}

func (s DummyCommonService) Update(commonId uint64, subdomain delivery.Common) error {
	index, err := storageIndexById(commonId)
	if err != nil {
		errUnknownCommonIdentifier := errors.New(fmt.Sprintf("Wrong delivery id %d", commonId))
		return errUnknownCommonIdentifier
	}

	subdomain.SetId(commonId)
	delivery.CommonStorage[index] = subdomain

	return nil
}

func (s DummyCommonService) Remove(commonId uint64) (bool, error) {
	index, err := storageIndexById(commonId)
	if err != nil {
		errUnknownCommonIdentifier := errors.New(fmt.Sprintf("Wrong delivery id %d", commonId))
		return false, errUnknownCommonIdentifier
	}

	delivery.CommonStorage = append(delivery.CommonStorage[:index], delivery.CommonStorage[index+1:]...)

	return true, nil
}

func storageIndexById(commonId uint64) (int, error) {
	for i, common := range delivery.CommonStorage {
		if common.Id() == commonId {
			return i, nil
		}
	}

	return -1, errors.New("id was not found")
}
