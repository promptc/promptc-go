package types

import (
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type FloatConstraint struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func (i *FloatConstraint) CanFit(v string) bool {
	i2, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return false
	}
	if i2 < i.Min {
		return false
	}
	if i2 > i.Max {
		return false
	}
	return true
}

var _ interfaces.Constraint = &FloatConstraint{}
