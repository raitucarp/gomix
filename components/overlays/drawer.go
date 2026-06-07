package overlays

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type drawer struct {
	component *components.Comp
}

func (c *drawer) Component() *components.Comp {
	return c.component
}

func (c *drawer) Element() *element.HtmlElement {
	return c.component.El
}

func (c *drawer) Children(children ...components.IsComponent) *drawer {
	c.component.Children(children...)
	return c
}

func Drawer(children ...components.IsComponent) *drawer {
	c := &drawer{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Drawer")
	c.component.El.AddAttribute("class", "chakra-drawer")

	c.component.Children(children...)
	return c
}
