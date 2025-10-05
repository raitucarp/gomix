package styles

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

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

type colorScale struct {
	value       uint
	transparent percentage
}

func (s *colorScale) Value() string {
	return fmt.Sprintf("%d", s.value)
}

func (s *colorScale) Transparent(p percentage) *colorScale {
	s.transparent = p
	return s
}

func ColorScale(value uint) *colorScale {
	return &colorScale{value: value}
}

type size struct {
	value any
}
