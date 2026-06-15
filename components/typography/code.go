package typography

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type code struct {
	component *components.Comp
}

func (c *code) Component() *components.Comp {
	return c.component
}

func (c *code) Element() *element.HtmlElement {
	return c.component.El
}

func (c *code) Children(children ...components.IsComponent) *code {
	c.component.Children(children...)
	return c
}

func Code(children ...components.IsComponent) *code {
	c := &code{
		component: &components.Comp{
			El: element.Code(element.Text("")).Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Code")
	c.component.El.AddAttribute("class", "chakra-code")

	c.component.Children(children...)
	return c
}
