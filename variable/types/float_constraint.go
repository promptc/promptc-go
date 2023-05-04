package types

import (
	"errors"
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type FloatConstraint struct {
	BaseConstraint
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

func (i *FloatConstraint) CanFit(v string) bool {
	i2, err := strconv.ParseFloat(v, 64)
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

func (i *FloatConstraint) String() string {
	return hjsonNoIndent(*i)
}

func (i *FloatConstraint) Validate() error {
	if i.Min != nil && i.Max != nil && *i.Min > *i.Max {
		return errors.New("min cannot be greater than max")
	}
	return nil
}

var _ interfaces.Constraint = &FloatConstraint{}
