package types

type NilConstraint struct {
}

func (c *NilConstraint) CanFit(value string) bool {
	return true
}

func (c *NilConstraint) String() string {
	return ""
}
