package layout

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type spacer struct {
	component *components.Comp
}

func (c *spacer) Component() *components.Comp {
	return c.component
}

func (c *spacer) Element() *element.HtmlElement {
	return c.component.El
}

func (c *spacer) Children(children ...components.IsComponent) *spacer {
	c.component.Children(children...)
	return c
}

func Spacer(children ...components.IsComponent) *spacer {
	c := &spacer{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Spacer")
	c.component.El.AddAttribute("class", "chakra-spacer")

	c.component.Children(children...)
	return c
}
