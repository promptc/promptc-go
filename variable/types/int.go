package types

import (
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type IntType struct {
	value      int64
	hasVal     bool
	name       string
	constraint interfaces.Constraint
}

func (i *IntType) Value() string {
	if i.hasVal {
		return strconv.FormatInt(i.value, 10)
	}
	return ""
}

func (i *IntType) SetValue(s string) bool {
	if !i.constraint.CanFit(s) {
		return false
	}
	i2, _ := strconv.ParseInt(s, 10, 64)
	i.value = i2
	i.hasVal = true
	return true
}

func (i *IntType) HasValue() bool {
	return i.hasVal
}

func (i *IntType) Type() string {
	return "int"
}

func (i *IntType) Name() string {
	return i.name
}

func (i *IntType) Constraint() interfaces.Constraint {
	return i.constraint
}

func (i *IntType) SetConstraint(c interfaces.Constraint) {
	i.constraint = c
}

var _ interfaces.Variable = &IntType{}

func NewInt(name string) *IntType {
	return &IntType{
		name: name,
	}
}
