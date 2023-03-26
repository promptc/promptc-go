package variable

type Variable interface {
	Type() string
	Name() string
	Value() string
	SetValue(string) bool
	HasValue() bool
}

type Constraint interface {
	CanFit(string) bool
}
