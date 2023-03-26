package types

import "github.com/promptc/promptc-go/variable"

type StringConstraint struct {
	MinLength int
	MaxLength int
}

func (s StringConstraint) CanFit(s2 string) bool {
	r := []rune(s2)
	if len(r) < s.MinLength {
		return false
	}
	if len(r) > s.MaxLength {
		return false
	}
	return true
}

var _ variable.Constraint = StringConstraint{}
