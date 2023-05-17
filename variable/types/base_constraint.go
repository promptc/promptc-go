package types

type BaseConstraint struct {
	Default     string `json:"default"`
	Description string `json:"description"`
}

func (b *BaseConstraint) DescriptionStr() *string {
	return &b.Description
}

func (b *BaseConstraint) ToMap() map[string]any {
	m := make(map[string]any)
	m["default"] = b.Default
	m["description"] = b.Description
	return m
}
