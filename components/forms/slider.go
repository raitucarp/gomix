package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type slider struct {
	component *components.Comp
}

func (c *slider) Component() *components.Comp {
	return c.component
}

func (c *slider) Element() *element.HtmlElement {
	return c.component.El
}

func (c *slider) Children(children ...components.IsComponent) *slider {
	c.component.Children(children...)
	return c
}

func Slider(children ...components.IsComponent) *slider {
	c := &slider{
		component: &components.Comp{
			El: element.Input().Element(),
		},
	}

	c.component.El.AddAttribute("type", "range")

	c.component.El.AddAttribute("data-gomix-component", "Slider")
	c.component.El.AddAttribute("class", "chakra-slider")

	c.component.Children(children...)
	return c
}
