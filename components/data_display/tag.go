package data_display

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type tag struct {
	component *components.Comp
}

func (c *tag) Component() *components.Comp {
	return c.component
}

func (c *tag) Element() *element.HtmlElement {
	return c.component.El
}

func (c *tag) Children(children ...components.IsComponent) *tag {
	c.component.Children(children...)
	return c
}

func Tag(children ...components.IsComponent) *tag {
	c := &tag{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Tag")
	c.component.El.AddAttribute("class", "chakra-tag")

	c.component.Children(children...)
	return c
}
