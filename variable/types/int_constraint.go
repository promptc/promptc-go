package types

import (
	"github.com/promptc/promptc-go/variable"
	"strconv"
)

type IntConstraint struct {
	Min int64
	Max int64
}

func (i IntConstraint) CanFit(v string) bool {
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

var _ variable.Constraint = IntConstraint{}
