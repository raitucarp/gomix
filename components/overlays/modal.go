package overlays

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type modal struct {
	component *components.Comp
}

func (c *modal) Component() *components.Comp {
	return c.component
}

func (c *modal) Element() *element.HtmlElement {
	return c.component.El
}

func (c *modal) Children(children ...components.IsComponent) *modal {
	c.component.Children(children...)
	return c
}

func Modal(children ...components.IsComponent) *modal {
	c := &modal{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Modal")
	c.component.El.AddAttribute("class", "chakra-modal")

	c.component.Children(children...)
	return c
}
