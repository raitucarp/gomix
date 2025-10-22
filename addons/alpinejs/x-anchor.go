package alpinejs

import "fmt"

type xAnchor struct {
	alpineRef *alpine
	ref       string
	position  string
	offset    string
	noStyle   bool
}

func (a *alpine) Anchor(ref string) *xAnchor {
	return &xAnchor{alpineRef: a, ref: "$refs." + ref}
}

func (x *xAnchor) Bottom() *xAnchor {
	x.position = "bottom"
	return x
}

func (x *xAnchor) BottomStart() *xAnchor {
	x.position = "bottom-start"
	return x
}

func (x *xAnchor) BottomEnd() *xAnchor {
	x.position = "bottom-end"
	return x
}

func (x *xAnchor) Top() *xAnchor {
	x.position = "top"
	return x
}

func (x *xAnchor) TopStart() *xAnchor {
	x.position = "top-start"
	return x
}

func (x *xAnchor) TopEnd() *xAnchor {
	x.position = "top-end"
	return x
}

func (x *xAnchor) Left() *xAnchor {
	x.position = "left"
	return x
}

func (x *xAnchor) LeftStart() *xAnchor {
	x.position = "left-start"
	return x
}

func (x *xAnchor) LeftEnd() *xAnchor {
	x.position = "left-end"
	return x
}

func (x *xAnchor) Right() *xAnchor {
	x.position = "right"
	return x
}

func (x *xAnchor) RightStart() *xAnchor {
	x.position = "right-start"
	return x
}

func (x *xAnchor) RighttEnd() *xAnchor {
	x.position = "right-end"
	return x
}

func (x *xAnchor) Offset(offset float64) *xAnchor {
	x.offset = fmt.Sprintf("%#v", offset)
	return x
}

func (x *xAnchor) NoStyle() *xAnchor {
	x.noStyle = true
	return x
}

func (x *xAnchor) AttachToID(id string) *xAnchor {
	x.ref = fmt.Sprintf("document.getElementById('%s')", id)
	return x
}

func (x *xAnchor) Attribute() *alpine {
	attributeName := "x-anchor"

	if x.position != "" {
		attributeName += "." + x.position
	}

	if x.offset != "" {
		attributeName += ".offset." + x.offset
	}

	if x.noStyle {
		attributeName += ".nostyle"
	}

	x.alpineRef.el.AddAttribute(attributeName, x.ref)
	return x.alpineRef
}
