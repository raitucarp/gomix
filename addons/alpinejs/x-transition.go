package alpinejs

import (
	"fmt"
	"strings"
)

type xTransition struct {
	alpineRef      *alpine
	directiveName  string
	directiveClass []string
	modifiers      []string
	value          string
}

func (a *alpine) Transition() *xTransition {
	return &xTransition{alpineRef: a}
}

func (x *xTransition) Enter(classNames ...string) *xTransition {
	x.directiveName = "enter"
	x.directiveClass = classNames
	return x
}

func (x *xTransition) EnterStart(classNames ...string) *xTransition {
	x.directiveName = "enter-start"
	x.directiveClass = classNames
	return x
}

func (x *xTransition) EnterEnd(classNames ...string) *xTransition {
	x.directiveName = "enter-end"
	x.directiveClass = classNames
	return x
}

func (x *xTransition) Leave(classNames ...string) *xTransition {
	x.directiveName = "leave"
	x.directiveClass = classNames
	return x
}

func (x *xTransition) LeaveStart(classNames ...string) *xTransition {
	x.directiveName = "leave-start"
	x.directiveClass = classNames
	return x
}

func (x *xTransition) LeaveEnd(classNames ...string) *xTransition {
	x.directiveName = "leave-end"
	x.directiveClass = classNames
	return x
}

func (x *xTransition) Modifiers(modifiers ...string) *xTransition {
	x.modifiers = append(x.modifiers, modifiers...)
	return x
}

func (x *xTransition) Attribute() *alpine {
	attributeName := "x-transition"
	if x.directiveName != "" {
		attributeName = fmt.Sprintf("x-transition:%s", x.directiveName)
		if len(x.directiveClass) > 0 {
			x.alpineRef.el.AddAttribute(attributeName, strings.Join(x.directiveClass, " "))
		} else {
			x.alpineRef.el.AddAttributeBool(attributeName)
		}

		return x.alpineRef
	}

	if len(x.modifiers) > 0 {
		a := []string{attributeName}
		a = append(a, x.modifiers...)
		attributeName = strings.Join(a, ".")
	}

	x.alpineRef.el.AddAttributeBool(attributeName)
	return x.alpineRef
}
