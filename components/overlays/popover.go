package overlays

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type popover struct {
	component *components.Comp
}

func (c *popover) Component() *components.Comp {
	return c.component
}

func (c *popover) Element() *element.HtmlElement {
	return c.component.El
}

func (c *popover) Children(children ...components.IsComponent) *popover {
	c.component.Children(children...)
	return c
}

func Popover(children ...components.IsComponent) *popover {
	c := &popover{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Popover")
	c.component.El.AddAttribute("class", "chakra-popover")

	c.component.Children(children...)
	return c
}
