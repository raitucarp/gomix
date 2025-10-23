package layout

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type stack struct {
	component *components.Comp
}

func (s *stack) Component() *components.Comp {
	return s.component
}

func (s *stack) Children(components ...components.IsComponent) *stack {
	s.component.Children(components...)
	return s
}

func Stack(children ...components.IsComponent) *stack {
	s := &stack{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	s.component.Flex().FlexCol().Gap(8).Children(children...)

	return s
}

func VStack(components ...components.IsComponent) *stack {
	s := Stack(components...)

	return s
}

func HStack(components ...components.IsComponent) *stack {
	s := Stack(components...)
	s.Component().FlexRow()
	return s
}

func (s *stack) Element() *element.HtmlElement {
	return s.component.El
}
