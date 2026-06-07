package feedback

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type toast struct {
	component *components.Comp
}

func (c *toast) Component() *components.Comp {
	return c.component
}

func (c *toast) Element() *element.HtmlElement {
	return c.component.El
}

func (c *toast) Children(children ...components.IsComponent) *toast {
	c.component.Children(children...)
	return c
}

func Toast(children ...components.IsComponent) *toast {
	c := &toast{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Toast")
	c.component.El.AddAttribute("class", "chakra-toast")

	c.component.Children(children...)
	return c
}
