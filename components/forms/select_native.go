package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type selectNative struct {
	component *components.Comp
}

func (c *selectNative) Component() *components.Comp {
	return c.component
}

func (c *selectNative) Element() *element.HtmlElement {
	return c.component.El
}

func (c *selectNative) Children(children ...components.IsComponent) *selectNative {
	c.component.Children(children...)
	return c
}

func SelectNative(children ...components.IsComponent) *selectNative {
	c := &selectNative{
		component: &components.Comp{
			El: element.Select().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "SelectNative")
	c.component.El.AddAttribute("class", "chakra-select-native")

	c.component.Children(children...)
	return c
}
