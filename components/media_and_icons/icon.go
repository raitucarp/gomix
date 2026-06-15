package media_and_icons

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type icon struct {
	component *components.Comp
}

func (c *icon) Component() *components.Comp {
	return c.component
}

func (c *icon) Element() *element.HtmlElement {
	return c.component.El
}

func (c *icon) Children(children ...components.IsComponent) *icon {
	c.component.Children(children...)
	return c
}

func Icon(children ...components.IsComponent) *icon {
	c := &icon{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Icon")
	c.component.El.AddAttribute("class", "chakra-icon")

	c.component.Children(children...)
	return c
}
