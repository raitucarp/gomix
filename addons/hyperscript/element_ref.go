package hyperscript

type elementRef struct {
	element string
}

func (e *elementRef) String() string {
	return e.element
}

func (e *elementRef) IsRef() {}

func Element(element string) *elementRef {
	return &elementRef{element: element}
}

func DocumentBody() *elementRef {
	return &elementRef{element: "document.body"}
}
