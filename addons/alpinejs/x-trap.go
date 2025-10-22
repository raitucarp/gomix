package alpinejs

type xTrap struct {
	alpineRef *alpine
	modifier  string
	value     string
}

func (a *alpine) Trap(value string) *xTrap {
	return &xTrap{alpineRef: a, value: value}
}

func (x *xTrap) Inert() *xTrap {
	x.modifier = "inert"
	return x
}

func (x *xTrap) NoScroll() *xTrap {
	x.modifier = "noscroll"
	return x
}

func (x *xTrap) NoReturn() *xTrap {
	x.modifier = "noreturn"
	return x
}

func (x *xTrap) NoAutofocus() *xTrap {
	x.modifier = "noautofocus"
	return x
}

func (x *xTrap) Attribute() *alpine {
	attributeName := "x-trap"
	if x.modifier != "" {
		attributeName = attributeName + "." + x.modifier

	}

	x.alpineRef.el.AddAttribute(attributeName, x.value)
	return x.alpineRef
}
