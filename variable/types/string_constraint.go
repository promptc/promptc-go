package types

import (
	"github.com/promptc/promptc-go/variable/interfaces"
)

type StringConstraint struct {
	MinLength int `json:"min_length"`
	MaxLength int `json:"max_length"`
}

func (s *StringConstraint) CanFit(s2 string) bool {
	r := []rune(s2)
	if len(r) < s.MinLength {
		return false
	}
	if len(r) > s.MaxLength {
		return false
	}
	return true
}

var _ interfaces.Constraint = &StringConstraint{}
