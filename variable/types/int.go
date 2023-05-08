package types

import (
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type IntType struct {
	BaseType
}

func (i *IntType) SetValue(s string) bool {
	if !i.constraint.CanFit(s) {
		return false
	}
	i2, _ := strconv.ParseInt(s, 10, 64)
	i.SetValueInternal(strconv.FormatInt(i2, 10), i2)
	return true
}

func (i *IntType) Type() string {
	return "int"
}

var _ interfaces.Variable = &IntType{}

func NewInt(name string) *IntType {
	return &IntType{
		BaseType: BaseType{
			name: name,
		},
	}
}
