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

type RatioValue struct {
	left  Value
	right Value
}

func (v *RatioValue) Value() string {
	return fmt.Sprintf("%s / %s", v.left.Value(), v.right.Value())
}

func Fraction(x1 Value, x2 Value) *RatioValue {
	v := &RatioValue{left: x1, right: x2}
	return v
}

func Ratio(x1 Value, x2 Value) *RatioValue {
	v := &RatioValue{left: x1, right: x2}
	return v
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
