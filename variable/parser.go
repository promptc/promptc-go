package variable

import (
	"errors"
	"fmt"
	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/variable/interfaces"
	"github.com/promptc/promptc-go/variable/types"
	"strings"
)

func Parse(singleLine string) (interfaces.Variable, error) {
	nameAndTail := strings.SplitN(singleLine, ":", 2)
	if len(nameAndTail) != 2 {
		return nil, fmt.Errorf("failed to parse variable %s", singleLine)
	}
	name := strings.TrimSpace(nameAndTail[0])
	return ParseKeyValue(name, nameAndTail[1])
}

func ParseKeyValue(name, tail string) (interfaces.Variable, error) {
	tail = strings.TrimSpace(tail)
	if tail == "" {
		v := types.NewString(name)
		v.SetConstraint(&types.NilConstraint{})
		return v, nil
	}
	typeAndTail := strings.SplitN(tail, "{", 2)
	vType := strings.TrimSpace(typeAndTail[0])
	cons := ""
	if len(typeAndTail) == 2 {
		cons = "{" + strings.TrimSpace(typeAndTail[1])
	}
	v := typeFactory(vType, name)
	if v == nil {
		return nil, errors.New("unknown type of " + vType + " for " + name)
	}

	if cons == "" {
		v.SetConstraint(&types.NilConstraint{})
	} else {
		_cons, err := consFactory(vType, cons)
		if err != nil {
			_cons = &types.NilConstraint{}
		}
		v.SetConstraint(_cons)
		return v, err
	}
	return v, nil
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

func consFactory(varType string, con string) (interfaces.Constraint, error) {
	var consA interfaces.Constraint
	switch varType {
	case "string":
		consA = &types.StringConstraint{}
	case "int":
		consA = &types.IntConstraint{}
	case "float":
		consA = &types.FloatConstraint{}
	default:
		return nil, errors.New("unknown type of " + varType)
	}
	if err := hjson.Unmarshal([]byte(con), consA); err != nil {
		return nil, err
	}
	return consA, nil
}
