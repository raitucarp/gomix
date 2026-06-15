package layout

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type simpleGrid struct {
	component *components.Comp
}

func (c *simpleGrid) Component() *components.Comp {
	return c.component
}

func (c *simpleGrid) Element() *element.HtmlElement {
	return c.component.El
}

func (c *simpleGrid) Children(children ...components.IsComponent) *simpleGrid {
	c.component.Children(children...)
	return c
}

func SimpleGrid(children ...components.IsComponent) *simpleGrid {
	c := &simpleGrid{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "SimpleGrid")
	c.component.El.AddAttribute("class", "chakra-simple-grid")

	c.component.Children(children...)
	return c
}
