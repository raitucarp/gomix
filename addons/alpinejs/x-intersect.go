package alpinejs

import (
	"fmt"
	"strings"
)

type xIntersect struct {
	alpineRef     *alpine
	directiveName string
	modifiers     []string
	value         string
}

func (a *alpine) Intersect(value string) *xIntersect {
	return &xIntersect{alpineRef: a, value: value}
}

func (x *xIntersect) Enter() *xIntersect {
	x.directiveName = "enter"
	return x
}

func (x *xIntersect) Leave() *xIntersect {
	x.directiveName = "leave"
	return x
}

func (x *xIntersect) Once() *xIntersect {
	x.modifiers = append(x.modifiers, "once")
	return x
}

func (x *xIntersect) Half() *xIntersect {
	x.modifiers = append(x.modifiers, "half")
	return x
}

func (x *xIntersect) Threshold(percentage float64) *xIntersect {
	x.modifiers = append(x.modifiers, "half")
	x.modifiers = append(x.modifiers, fmt.Sprintf("%#v", percentage))
	return x
}

func (x *xIntersect) Margin(values ...string) *xIntersect {
	x.modifiers = append(x.modifiers, "margin")
	x.modifiers = append(x.modifiers, values...)
	return x
}

func (x *xIntersect) Attribute() *alpine {
	attributeName := "x-intersect"

	if x.directiveName != "" {
		attributeName = fmt.Sprintf("x-intersect:%s", x.directiveName)
	}

	if len(x.modifiers) > 0 {
		attributeName = fmt.Sprintf("%s.%s", attributeName, strings.Join(x.modifiers, "."))
	}

	x.alpineRef.el.AddAttribute(attributeName, x.value)

	return x.alpineRef
}
