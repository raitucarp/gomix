package overlays

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type menu struct {
	component *components.Comp
}

func (c *menu) Component() *components.Comp {
	return c.component
}

func (c *menu) Element() *element.HtmlElement {
	return c.component.El
}

func (c *menu) Children(children ...components.IsComponent) *menu {
	c.component.Children(children...)
	return c
}

func Menu(children ...components.IsComponent) *menu {
	c := &menu{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Menu")
	c.component.El.AddAttribute("class", "chakra-menu")

	c.component.Children(children...)
	return c
}
