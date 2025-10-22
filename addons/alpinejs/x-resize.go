package alpinejs

import (
	"fmt"
	"strings"
)

type xResize struct {
	alpineRef *alpine
	document  bool
	value     []string
}

func (a *alpine) Resize() *xResize {
	return &xResize{alpineRef: a}
}

func (x *xResize) AssignTo(varName string, to string) *xResize {
	x.value = append(x.value, fmt.Sprintf("$%s: %s", varName, to))
	return x
}

func (x *xResize) Document() *xResize {
	x.document = true
	return x
}

func (x *xResize) Attribute() *alpine {
	attributeName := "x-resize"
	if x.document {
		attributeName += ":document"
	}
	x.alpineRef.el.AddAttribute(attributeName, strings.Join(x.value, "; "))
	return x.alpineRef
}
