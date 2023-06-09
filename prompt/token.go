package prompt

type BlockToken struct {
	Text string         `json:"text"`
	Kind BlockTokenKind `json:"kind"`
}

type BlockTokenKind string

const BlockTokenKindLiter BlockTokenKind = "liter"
const BlockTokenKindVar BlockTokenKind = "var"
const BlockTokenKindScript BlockTokenKind = "script"
const BlockTokenKindReservedQuota BlockTokenKind = "reserved_quota"

func (b BlockTokenKind) String() string {
	return string(b)
}
