package data_display

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type table struct {
	component *components.Comp
}

func (c *table) Component() *components.Comp {
	return c.component
}

func (c *table) Element() *element.HtmlElement {
	return c.component.El
}

func (c *table) Children(children ...components.IsComponent) *table {
	c.component.Children(children...)
	return c
}

func Table(children ...components.IsComponent) *table {
	c := &table{
		component: &components.Comp{
			El: element.Table().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Table")
	c.component.El.AddAttribute("class", "chakra-table")

	c.component.Children(children...)
	return c
}
