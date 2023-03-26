package types

import (
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type IntConstraint struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

func (i *IntConstraint) CanFit(v string) bool {
	i2, err := strconv.ParseInt(v, 10, 64)
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

var _ interfaces.Constraint = &IntConstraint{}
