package types

import (
	"github.com/promptc/promptc-go/variable/interfaces"
)

type StringType struct {
	name       string
	constraint interfaces.Constraint
	value      string
	hasVal     bool
}

func (s *StringType) Value() string {
	return s.value
}

func (s *StringType) SetValue(s2 string) bool {
	if s.constraint.CanFit(s2) {
		s.value = s2
		s.hasVal = true
		return true
	}
	return false
}

func (s *StringType) HasValue() bool {
	return s.hasVal
}

func (s *StringType) Type() string {
	return "string"
}

func (s *StringType) Name() string {
	return s.name
}

func (s *StringType) Constraint() interfaces.Constraint {
	return s.constraint
}

func (s *StringType) SetConstraint(c interfaces.Constraint) {
	s.constraint = c
}

func (s *StringType) Description() string {
	descri := s.constraint.DescriptionStr()
	if descri != nil {
		return *descri
	}
	return s.name
}

var _ interfaces.Variable = &StringType{}

func NewString(name string) *StringType {
	return &StringType{
		name: name,
	}
}
