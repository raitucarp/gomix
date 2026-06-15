package data_display

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type badge struct {
	component *components.Comp
}

func (c *badge) Component() *components.Comp {
	return c.component
}

func (c *badge) Element() *element.HtmlElement {
	return c.component.El
}

func (c *badge) Children(children ...components.IsComponent) *badge {
	c.component.Children(children...)
	return c
}

func Badge(children ...components.IsComponent) *badge {
	c := &badge{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Badge")
	c.component.El.AddAttribute("class", "chakra-badge")

	c.component.Children(children...)
	return c
}
