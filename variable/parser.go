package variable

import (
	"fmt"
	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/variable/interfaces"
	"github.com/promptc/promptc-go/variable/types"
	"strings"
)

func Parse(singleLine string) interfaces.Variable {
	nameAndTail := strings.SplitN(singleLine, ":", 2)
	if len(nameAndTail) != 2 {
		return nil
	}
	name := strings.TrimSpace(nameAndTail[0])
	return ParseKeyValue(name, nameAndTail[1])
}

func ParseKeyValue(name, tail string) interfaces.Variable {
	typeAndTail := strings.SplitN(tail, "{", 2)
	vType := strings.TrimSpace(typeAndTail[0])
	cons := ""
	if len(typeAndTail) == 2 {
		cons = "{" + strings.TrimSpace(typeAndTail[1])
	}
	v := typeFactory(vType, name)
	if v == nil {
		return nil
	}
	if cons == "" {
		v.SetConstraint(&types.NilConstraint{})
	} else {
		v.SetConstraint(consFactory(vType, cons))
	}
	return v
}

func typeFactory(varType string, name string) interfaces.Variable {
	switch varType {
	case "string":
		return types.NewString(name)
	case "int":
		return types.NewInt(name)
	case "float":
		return types.NewFloat(name)
	default:
		return nil
	}
}

func consFactory(varType string, con string) interfaces.Constraint {
	var consA interfaces.Constraint
	switch varType {
	case "string":
		consA = &types.StringConstraint{}
	case "int":
		consA = &types.IntConstraint{}
	case "float":
		consA = &types.FloatConstraint{}
	default:
		return nil
	}
	if err := hjson.Unmarshal([]byte(con), consA); err != nil {
		fmt.Println("Failed to parse constraint", con)
		panic(err)
	}
	return consA
}
