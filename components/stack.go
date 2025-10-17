package components

import (
	"github.com/raitucarp/gomix/element"
)

type stack struct {
	component *component
}

func (s *stack) Component() *component {
	return s.component
}

func (s *stack) Children(components ...IsComponent) *stack {
	s.component.Children(components...)
	return s
}

func Stack(components ...IsComponent) *stack {
	s := &stack{
		component: &component{
			el: div().Element(),
		},
	}

	s.component.Flex().FlexCol().Gap(8).Children(components...)

	return s
}

func VStack(components ...IsComponent) *stack {
	s := Stack(components...)

	return s
}

func HStack(components ...IsComponent) *stack {
	s := Stack(components...)
	s.Component().FlexRow()
	return s
}

func (s *stack) Element() *element.HtmlElement {
	return s.component.el
}
