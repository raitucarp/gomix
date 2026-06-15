package data_display

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type divider struct {
	component *components.Comp
}

func (c *divider) Component() *components.Comp {
	return c.component
}

func (c *divider) Element() *element.HtmlElement {
	return c.component.El
}

func (c *divider) Children(children ...components.IsComponent) *divider {
	c.component.Children(children...)
	return c
}

func Divider(children ...components.IsComponent) *divider {
	c := &divider{
		component: &components.Comp{
			El: element.Hr().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Divider")
	c.component.El.AddAttribute("class", "chakra-divider")

	c.component.Children(children...)
	return c
}
