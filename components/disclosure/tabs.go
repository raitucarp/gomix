package disclosure

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type tabs struct {
	component *components.Comp
}

func (c *tabs) Component() *components.Comp {
	return c.component
}

func (c *tabs) Element() *element.HtmlElement {
	return c.component.El
}

func (c *tabs) Children(children ...components.IsComponent) *tabs {
	c.component.Children(children...)
	return c
}

func Tabs(children ...components.IsComponent) *tabs {
	c := &tabs{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Tabs")
	c.component.El.AddAttribute("class", "chakra-tabs")

	c.component.Children(children...)
	return c
}
