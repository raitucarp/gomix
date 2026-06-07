package typography

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type heading struct {
	component *components.Comp
}

func (c *heading) Component() *components.Comp {
	return c.component
}

func (c *heading) Element() *element.HtmlElement {
	return c.component.El
}

func (c *heading) Children(children ...components.IsComponent) *heading {
	c.component.Children(children...)
	return c
}

func Heading(children ...components.IsComponent) *heading {
	c := &heading{
		component: &components.Comp{
			El: element.H1(element.Text("")).Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Heading")
	c.component.El.AddAttribute("class", "chakra-heading")

	c.component.Children(children...)
	return c
}
