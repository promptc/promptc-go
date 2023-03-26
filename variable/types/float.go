package types

import (
	"fmt"
	"github.com/promptc/promptc-go/variable"
	"strconv"
)

type FloatType struct {
	value      float64
	hasVal     bool
	name       string
	constraint FloatConstraint
}

func (i FloatType) Value() string {
	if i.hasVal {
		return fmt.Sprintf("%.2f", i.value)
	}
	return ""
}

func (i FloatType) SetValue(s string) bool {
	if !i.constraint.CanFit(s) {
		return false
	}
	i2, _ := strconv.ParseFloat(s, 64)
	i.value = i2
	i.hasVal = true
	return true
}

func (i FloatType) HasValue() bool {
	return i.hasVal
}

func (i FloatType) Type() string {
	return "float"
}

func (i FloatType) Name() string {
	return i.name
}

var _ variable.Variable = FloatType{}
