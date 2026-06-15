package disclosure

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type visuallyHidden struct {
	component *components.Comp
}

func (c *visuallyHidden) Component() *components.Comp {
	return c.component
}

func (c *visuallyHidden) Element() *element.HtmlElement {
	return c.component.El
}

func (c *visuallyHidden) Children(children ...components.IsComponent) *visuallyHidden {
	c.component.Children(children...)
	return c
}

func VisuallyHidden(children ...components.IsComponent) *visuallyHidden {
	c := &visuallyHidden{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "VisuallyHidden")
	c.component.El.AddAttribute("class", "chakra-visually-hidden")

	c.component.Children(children...)
	return c
}
