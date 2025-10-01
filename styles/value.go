package styles

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

type customProperty struct {
	name         string
	inherits     bool
	initialValue any
}

type customVariableProp struct {
	name string
}

func CustomVar(name string) *customVariableProp {
	return &customVariableProp{name: name}
}

func (v *customVariableProp) Value() string {
	return fmt.Sprintf("var(--%s)", strcase.ToDelimited(v.name, '-'))
}

type customValue interface {
	Value() string
}

type val struct {
	value any
}

func (v *val) Value() string {
	switch val := v.value.(type) {
	case string:
		return val
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", v.value)
	}
}

func Val(value any) *val {
	v := val{value: value}
	return &v
}

type scale struct {
	value       uint
	transparent percentage
}

func (s *scale) Value() string {
	return fmt.Sprintf("%d", s.value)
}

func (s *scale) Transparent(p percentage) *scale {
	s.transparent = p
	return s
}

func Scale(value uint) *scale {
	return &scale{value: value}
}

type size struct {
	value any
}

