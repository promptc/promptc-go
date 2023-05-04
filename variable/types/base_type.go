package types

import "github.com/promptc/promptc-go/variable/interfaces"

type BaseType struct {
	hasVal     bool
	name       string
	constraint interfaces.Constraint
	valStr     string
	rawVal     any
}

func (i *BaseType) HasValue() bool {
	return i.hasVal
}

func (i *BaseType) Name() string {
	return i.name
}

func (i *BaseType) Description() string {
	descri := i.constraint.DescriptionStr()
	if descri != nil {
		return *descri
	}
	return i.name
}

func (i *BaseType) Constraint() interfaces.Constraint {
	return i.constraint
}

func (i *BaseType) SetConstraint(c interfaces.Constraint) {
	i.constraint = c
}

func (i *BaseType) Value() string {
	if i.hasVal {
		return i.valStr
	}
	return ""
}

func (i *BaseType) SetValueInternal(strVal string, raw any) {
	i.valStr = strVal
	i.rawVal = raw
	i.hasVal = true
}
