package feedback

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type alert struct {
	component *components.Comp
}

func (c *alert) Component() *components.Comp {
	return c.component
}

func (c *alert) Element() *element.HtmlElement {
	return c.component.El
}

func (c *alert) Children(children ...components.IsComponent) *alert {
	c.component.Children(children...)
	return c
}

func Alert(children ...components.IsComponent) *alert {
	c := &alert{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Alert")
	c.component.El.AddAttribute("class", "chakra-alert")

	c.component.Children(children...)
	return c
}
