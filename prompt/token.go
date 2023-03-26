package prompt

type BlockToken struct {
	Text string
	Kind BlockTokenKind
}

type BlockTokenKind string

const BlockTokenKindLiter BlockTokenKind = "liter"
const BlockTokenKindVar BlockTokenKind = "var"
const BlockTokenKindScript BlockTokenKind = "script"

func (b BlockTokenKind) String() string {
	return string(b)
}
