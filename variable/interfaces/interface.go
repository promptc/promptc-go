package interfaces

type Variable interface {
	Type() string
	Name() string
	Value() string
	SetValue(string) bool
	HasValue() bool
	Constraint() Constraint
	SetConstraint(Constraint)
}

type Constraint interface {
	CanFit(string) bool
	String() string
}
