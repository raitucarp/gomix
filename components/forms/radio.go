package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type radio struct {
	component *components.Comp
}

func (c *radio) Component() *components.Comp {
	return c.component
}

func (c *radio) Element() *element.HtmlElement {
	return c.component.El
}

func (c *radio) Children(children ...components.IsComponent) *radio {
	c.component.Children(children...)
	return c
}

func Radio(children ...components.IsComponent) *radio {
	c := &radio{
		component: &components.Comp{
			El: element.Input().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Radio")
	c.component.El.AddAttribute("class", "chakra-radio")
	c.component.El.AddAttribute("type", "radio")

	c.component.Children(children...)
	return c
}

func (c *radio) ColorScheme(color string) *radio {
	c.component.El.AddAttribute("data-color-scheme", color)
	return c
}
func (c *radio) Size(s string) *radio {
	c.component.El.AddAttribute("data-size", s)
	return c
}
