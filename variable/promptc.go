package variable

import (
	"github.com/promptc/promptc-go/variable/interfaces"
	"strings"
)

func ToPromptcString(v interfaces.Variable) string {
	sb := strings.Builder{}
	sb.WriteString(v.Name())
	sb.WriteString(": ")
	sb.WriteString(v.Type())
	if v.Constraint() != nil {
		sb.WriteString(v.Constraint().String())
	}
	return sb.String()
}

func ToPromptcValue(v interfaces.Variable) string {
	sb := strings.Builder{}
	sb.WriteString(v.Type())
	if v.Constraint() != nil {
		sb.WriteString(v.Constraint().String())
	}
	return sb.String()
}
