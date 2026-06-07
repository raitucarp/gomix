package layout

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type grid struct {
	component *components.Comp
}

func (c *grid) Component() *components.Comp {
	return c.component
}

func (c *grid) Element() *element.HtmlElement {
	return c.component.El
}

func (c *grid) Children(children ...components.IsComponent) *grid {
	c.component.Children(children...)
	return c
}

func Grid(children ...components.IsComponent) *grid {
	c := &grid{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Grid")
	c.component.El.AddAttribute("class", "chakra-grid")

	c.component.Children(children...)
	return c
}
