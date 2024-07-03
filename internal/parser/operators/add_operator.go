package operators

import (
	"cheese-lang/internal/parser"
	"errors"
)

func AddOperator(v1 parser.VariableContainer, v2 parser.VariableContainer) (parser.VariableContainer, error) {
	value1, ok1 := v1.(*parser.ParmesanVariable)
	value2, ok2 := v2.(*parser.ParmesanVariable)
	if ok1 && ok2 {
		return &parser.ParmesanVariable{Value: value1.Value + value2.Value}, nil
	}

	value3, ok1 := v1.(*parser.GorgonzolaVariable)
	value4, ok2 := v2.(*parser.GorgonzolaVariable)
	if ok1 && ok2 {
		return &parser.GorgonzolaVariable{Value: value3.Value + value4.Value}, nil
	}

	return nil, errors.New("Add operator unsupported for types: " + v1.GetVariableType().String() + ", " + v2.GetVariableType().String())
}
