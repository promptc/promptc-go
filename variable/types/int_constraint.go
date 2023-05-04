package types

import (
	"errors"
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type IntConstraint struct {
	Min         *int64  `json:"min,omitempty"`
	Max         *int64  `json:"max,omitempty"`
	Default     string  `json:"default"`
	Description *string `json:"description"`
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
	return hjsonNoIndent(*i)
}

func (i *IntConstraint) Validate() error {
	if i.Min != nil && i.Max != nil && *i.Min > *i.Max {
		return errors.New("min cannot be greater than max")
	}
	return nil
}

func (i *IntConstraint) DescriptionStr() *string {
	return i.Description
}

var _ interfaces.Constraint = &IntConstraint{}
