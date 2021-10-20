package coupon

import (
	"fmt"
)

type Coupon struct {
	Code string
	Percent uint64
}

var couponList []Coupon

func (c *Coupon) String() string {
	return fmt.Sprintf("Coupon %s - %d%s", c.Code, c.Percent, "%")
}

