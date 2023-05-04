package types

import (
	"github.com/promptc/promptc-go/variable/interfaces"
)

type StringType struct {
	BaseType
	name       string
	constraint interfaces.Constraint
	value      string
	hasVal     bool
}

func (s *StringType) SetValue(s2 string) bool {
	if s.constraint.CanFit(s2) {
		s.SetValueInternal(s2, s2)
		return true
	}
	return false
}

func (s *StringType) Type() string {
	return "string"
}

var _ interfaces.Variable = &StringType{}

func NewString(name string) *StringType {
	return &StringType{
		BaseType: BaseType{
			name: name,
		},
	}
}
