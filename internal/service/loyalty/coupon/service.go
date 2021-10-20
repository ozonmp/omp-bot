package coupon

import (
	"fmt"
)

type CouponService interface {
	Describe(CouponID uint64) (*Coupon, error)
	List(cursor uint64, limit uint64) ([]Coupon, error)
	Create(Coupon) (uint64, error)
	Update(CouponID uint64, coupon Coupon) error
	Remove(CouponID uint64) (bool, error)
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(cursor uint64, limit uint64) ([]Coupon, error) {
	if len(couponList) == 0 {
		return nil, nil
	}
	if int(cursor) >= len(couponList) {
		return nil, fmt.Errorf("coupon index is wrong")
	}
	if int(cursor + limit) >= len(couponList) {
		return couponList[cursor :], nil
	} else {
		return couponList[cursor : cursor+limit], nil
	}
}

func (s *Service) Describe(couponID uint64) (*Coupon, error) {
	if couponID < 0 || int(couponID) >= len(couponList) {
		return nil, fmt.Errorf("coupon index is wrong")
	}
	return &couponList[couponID], nil
}

func (s *Service) Create(coupon Coupon) (uint64, error) {
	if coupon.Percent < 0 {
		return 0, fmt.Errorf(": need positive percent value (%d)", coupon.Percent)
	}
	couponList = append(couponList, coupon)
	return uint64(len(couponList)), nil
}

func (s *Service) Remove (id uint64) (bool, error) {
	if id < 0 || int(id) >= len(couponList) {
		return false, fmt.Errorf(": id range check error (%d)", id)
	}
	couponList = append(couponList[:id], couponList[id+1:]...)
	return true, nil
}

func (s *Service) Update(CouponID uint64, coupon Coupon) error {
	if CouponID < 0 || int(CouponID) >= len(couponList) {
		return fmt.Errorf(": id range check error (%d)", CouponID)
	}
	couponList = append(couponList[:CouponID], couponList[CouponID+1:]...)
	return nil
}

