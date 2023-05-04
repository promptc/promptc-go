package types

type BaseConstraint struct {
	Default     string `json:"default"`
	Description string `json:"description"`
}

func (b *BaseConstraint) DescriptionStr() *string {
	return &b.Description
}
