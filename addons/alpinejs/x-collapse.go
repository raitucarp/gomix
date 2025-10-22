package alpinejs

import "strings"

type xCollapse struct {
	alpineRef *alpine
	modifier  string
}

func (a *alpine) Collapse() *xCollapse {
	return &xCollapse{alpineRef: a}
}

func (x *xCollapse) Duration(duration string) *xCollapse {
	x.modifier = strings.Join([]string{"duration", duration}, ".")
	return x
}

func (x *xCollapse) Min(value string) *xCollapse {
	x.modifier = strings.Join([]string{"min", value}, ".")
	return x
}

func (x *xCollapse) Attribute() *alpine {
	attributeName := "x-collapse"

	if x.modifier != "" {
		attributeName = attributeName + "." + x.modifier
	}

	x.alpineRef.el.AddAttributeBool(attributeName)
	return x.alpineRef
}
