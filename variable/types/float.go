package types

import (
	"fmt"
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type FloatType struct {
	value      float64
	hasVal     bool
	name       string
	constraint interfaces.Constraint
}

func (i *FloatType) Value() string {
	if i.hasVal {
		return fmt.Sprintf("%.2f", i.value)
	}
	return ""
}

func (i *FloatType) SetValue(s string) bool {
	if !i.constraint.CanFit(s) {
		return false
	}
	i2, _ := strconv.ParseFloat(s, 64)
	i.value = i2
	i.hasVal = true
	return true
}

func (i *FloatType) HasValue() bool {
	return i.hasVal
}

func (i *FloatType) Type() string {
	return "float"
}

func (i *FloatType) Name() string {
	return i.name
}

func (i *FloatType) Constraint() interfaces.Constraint {
	return i.constraint
}

func (i *FloatType) SetConstraint(c interfaces.Constraint) {
	i.constraint = c
}

func (i *FloatType) String() string {
	v := fmt.Sprintf("%s : %s\n", i.name, i.Type())
	if i.constraint != nil {
		v += fmt.Sprintf("%#v\n", i.constraint)
	}
	return v
}

func (i *FloatType) Description() string {
	descri := i.constraint.DescriptionStr()
	if descri != nil {
		return *descri
	}
	return i.name
}

var _ interfaces.Variable = &FloatType{}

func NewFloat(name string) *FloatType {
	return &FloatType{
		name: name,
	}
}
