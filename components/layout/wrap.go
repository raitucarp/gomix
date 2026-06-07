package layout

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type wrap struct {
	component *components.Comp
}

func (c *wrap) Component() *components.Comp {
	return c.component
}

func (c *wrap) Element() *element.HtmlElement {
	return c.component.El
}

func (c *wrap) Children(children ...components.IsComponent) *wrap {
	c.component.Children(children...)
	return c
}

func Wrap(children ...components.IsComponent) *wrap {
	c := &wrap{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Wrap")
	c.component.El.AddAttribute("class", "chakra-wrap")

	c.component.Children(children...)
	return c
}
