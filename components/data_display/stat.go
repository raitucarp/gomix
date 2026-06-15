package data_display

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type stat struct {
	component *components.Comp
}

func (c *stat) Component() *components.Comp {
	return c.component
}

func (c *stat) Element() *element.HtmlElement {
	return c.component.El
}

func (c *stat) Children(children ...components.IsComponent) *stat {
	c.component.Children(children...)
	return c
}

func Stat(children ...components.IsComponent) *stat {
	c := &stat{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Stat")
	c.component.El.AddAttribute("class", "chakra-stat")

	c.component.Children(children...)
	return c
}
