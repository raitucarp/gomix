package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type numberInput struct {
	component *components.Comp
}

func (c *numberInput) Component() *components.Comp {
	return c.component
}

func (c *numberInput) Element() *element.HtmlElement {
	return c.component.El
}

func (c *numberInput) Children(children ...components.IsComponent) *numberInput {
	c.component.Children(children...)
	return c
}

func NumberInput(children ...components.IsComponent) *numberInput {
	c := &numberInput{
		component: &components.Comp{
			El: element.Input().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "NumberInput")
	c.component.El.AddAttribute("class", "chakra-number-input")

	c.component.Children(children...)
	return c
}
