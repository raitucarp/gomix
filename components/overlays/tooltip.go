package overlays

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type tooltip struct {
	component *components.Comp
}

func (c *tooltip) Component() *components.Comp {
	return c.component
}

func (c *tooltip) Element() *element.HtmlElement {
	return c.component.El
}

func (c *tooltip) Children(children ...components.IsComponent) *tooltip {
	c.component.Children(children...)
	return c
}

func Tooltip(children ...components.IsComponent) *tooltip {
	c := &tooltip{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Tooltip")
	c.component.El.AddAttribute("class", "chakra-tooltip")

	c.component.Children(children...)
	return c
}
