package types

import "github.com/promptc/promptc-go/variable/interfaces"

type NilConstraint struct {
}

func (c *NilConstraint) CanFit(value string) bool {
	return true
}

func (c *NilConstraint) String() string {
	return ""
}

func (c *NilConstraint) Validate() error {
	return nil
}

func (c *NilConstraint) DescriptionStr() *string {
	return nil
}

func (c *NilConstraint) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

var _ interfaces.Constraint = &NilConstraint{}
