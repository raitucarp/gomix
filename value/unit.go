package value

import "fmt"

type unit[T valConstraint] struct {
	value  T
	suffix string
}

func (u *unit[T]) Value() string {
	return fmt.Sprintf("%#v%s", u.value, u.suffix)
}

func Rem[T valConstraint](value T) *unit[T] {
	return &unit[T]{value: value, suffix: "rem"}
}

func Em[T valConstraint](value T) *unit[T] {
	return &unit[T]{value: value, suffix: "em"}
}

func Px[T valConstraint](value T) *unit[T] {
	return &unit[T]{value: value, suffix: "px"}
}

func Percent(value float64) *unit[float64] {
	return &unit[float64]{value: value, suffix: "%"}
}

func Pi() *unit[string] {
	return &unit[string]{value: "pi", suffix: ""}
}
