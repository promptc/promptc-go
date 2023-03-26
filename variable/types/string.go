package types

import "github.com/promptc/promptc-go/variable"

type StringType struct {
	name       string
	constraint StringConstraint
	value      string
	hasVal     bool
}

func (s StringType) Value() string {
	return s.value
}

func (s StringType) SetValue(s2 string) bool {
	if s.constraint.CanFit(s2) {
		s.value = s2
		s.hasVal = true
		return true
	}
	return false
}

func (s StringType) HasValue() bool {
	return s.hasVal
}

func (s StringType) Type() string {
	return "string"
}

func (s StringType) Name() string {
	return s.name
}

var _ variable.Variable = StringType{}
