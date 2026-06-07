package layout

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type center struct {
	component *components.Comp
}

func (c *center) Component() *components.Comp {
	return c.component
}

func (c *center) Element() *element.HtmlElement {
	return c.component.El
}

func (c *center) Children(children ...components.IsComponent) *center {
	c.component.Children(children...)
	return c
}

func Center(children ...components.IsComponent) *center {
	c := &center{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Center")
	c.component.El.AddAttribute("class", "chakra-center")

	c.component.Children(children...)
	return c
}
