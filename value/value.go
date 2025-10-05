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

type literal struct {
	value string
}

func (l *literal) Value() string {
	return l.value
}

func Number[T valConstraint](val T) *number[T] {
	return &number[T]{value: val}
}

func Literal(v string) *literal {
	return &literal{value: v}
}

type variable struct {
	name string
}

func (v *variable) Value() string {
	return fmt.Sprintf("var(--%s)", v.name)
}

func Var(name string) *variable {
	return &variable{name: name}
}
