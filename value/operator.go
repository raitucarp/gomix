package value

import "fmt"

type operator struct {
	values []Value
	format string
}

func (op *operator) Value() string {
	var params []any
	for _, value := range op.values {
		params = append(params, value.Value())
	}

	return fmt.Sprintf(op.format, params...)
}

func Mul(x1 Value, x2 Value) *operator {
	op := &operator{format: "%s * %s", values: []Value{x1, x2}}
	return op
}

func Fraction(x1 Value, x2 Value) *operator {
	op := &operator{format: "%s / %s", values: []Value{x1, x2}}
	return op
}

func Add(x1 Value, x2 Value) *operator {
	op := &operator{format: "%s + %s", values: []Value{x1, x2}}
	return op
}

func Subtract(x1 Value, x2 Value) *operator {
	op := &operator{format: "%s - %s", values: []Value{x1, x2}}
	return op
}

func Neg(x1 Value) *operator {
	op := &operator{format: "-%s", values: []Value{x1}}
	return op
}
