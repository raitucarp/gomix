package alpinejs

type xMask struct {
	alpineRef *alpine
	dynamic   bool
	value     string
}

func (a *alpine) Mask(value string) *xMask {
	return &xMask{alpineRef: a, value: value}
}

func (x *xMask) Dynamic() *xMask {
	x.dynamic = true
	return x
}

func (x *xMask) Attribute() *alpine {
	if x.dynamic {
		x.alpineRef.el.AddAttribute("x-mask:dynamic", x.value)
	} else {
		x.alpineRef.el.AddAttribute("x-mask", x.value)
	}

	return x.alpineRef
}
