package types

import (
	"fmt"
	"github.com/promptc/promptc-go/variable/interfaces"
	"strconv"
)

type FloatType struct {
	BaseType
	value float64
}

func (i *FloatType) SetValue(s string) bool {
	if !i.constraint.CanFit(s) {
		return false
	}
	i2, _ := strconv.ParseFloat(s, 64)
	i.SetValueInternal(fmt.Sprintf("%.2f", i2), i2)
	return true
}

func (i *FloatType) Type() string {
	return "float"
}

func (i *FloatType) String() string {
	v := fmt.Sprintf("%s : %s\n", i.name, i.Type())
	if i.constraint != nil {
		v += fmt.Sprintf("%#v\n", i.constraint)
	}
	return v
}

var _ interfaces.Variable = &FloatType{}

func NewFloat(name string) *FloatType {
	return &FloatType{
		BaseType: BaseType{
			name: name,
		},
	}
}
