package types

import (
	"errors"
	"github.com/promptc/promptc-go/variable/interfaces"
)

type StringConstraint struct {
	BaseConstraint
	MinLength *int `json:"minLen"`
	MaxLength *int `json:"maxLen"`
}

func (s *StringConstraint) CanFit(s2 string) bool {
	r := []rune(s2)
	if s.MinLength != nil && len(r) < *s.MinLength {
		return false
	}
	if s.MaxLength != nil && len(r) > *s.MaxLength {
		return false
	}
	return true
}

func (s *StringConstraint) String() string {
	return hjsonNoIndent(*s)
}

func (s *StringConstraint) Validate() error {
	if s.MinLength != nil && s.MaxLength != nil && *s.MinLength > *s.MaxLength {
		return errors.New("min length cannot be greater than max length")
	}
	return nil
}

var _ interfaces.Constraint = &StringConstraint{}
