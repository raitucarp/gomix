package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type pinInput struct {
	component *components.Comp
}

func (c *pinInput) Component() *components.Comp {
	return c.component
}

func (c *pinInput) Element() *element.HtmlElement {
	return c.component.El
}

func (c *pinInput) Children(children ...components.IsComponent) *pinInput {
	c.component.Children(children...)
	return c
}

func PinInput(children ...components.IsComponent) *pinInput {
	c := &pinInput{
		component: &components.Comp{
			El: element.Input().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "PinInput")
	c.component.El.AddAttribute("class", "chakra-pin-input")

	c.component.Children(children...)
	return c
}
