package value

import "fmt"

type Value interface {
	Value() string
}

type typ interface {
	IsType()
}

type valConstraint interface {
	int | float32 | float64 | string
}

type number[T valConstraint] struct {
	value T
}

func (n *number[T]) Value() string {
	return fmt.Sprintf("%#v", n.value)
}

type LiteralValue struct {
	value string
}

func (l *LiteralValue) Value() string {
	return l.value
}

func Number[T valConstraint](val T) *number[T] {
	return &number[T]{value: val}
}

func Literal(v string) *LiteralValue {
	return &LiteralValue{value: v}
}

type CustomProperty struct {
	name string
}

func (p *CustomProperty) Value() string {
	return fmt.Sprintf("var(--%s)", p.name)
}

func Property(name string) *CustomProperty {
	return &CustomProperty{name: name}
}
