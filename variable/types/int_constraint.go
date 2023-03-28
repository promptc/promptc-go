package types

import (
	"github.com/promptc/promptc-go/utils"
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type IntConstraint struct {
	Min     *int64 `json:"min,omitempty"`
	Max     *int64 `json:"max,omitempty"`
	Default string `json:"default"`
}

func (i *IntConstraint) CanFit(v string) bool {
	i2, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return false
	}
	if i.Min != nil && i2 < *i.Min {
		return false
	}
	if i.Max != nil && i2 > *i.Max {
		return false
	}
	return true
}

func (i *IntConstraint) String() string {
	return utils.Hjson(*i)
}

var _ interfaces.Constraint = &IntConstraint{}
