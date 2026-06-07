package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type switchCmp struct {
	component *components.Comp
}

func (c *switchCmp) Component() *components.Comp {
	return c.component
}

func (c *switchCmp) Element() *element.HtmlElement {
	return c.component.El
}

func (c *switchCmp) Children(children ...components.IsComponent) *switchCmp {
	c.component.Children(children...)
	return c
}

func SwitchCmp(children ...components.IsComponent) *switchCmp {
	c := &switchCmp{
		component: &components.Comp{
			El: element.Input().Element(),
		},
	}

	c.component.El.AddAttribute("type", "checkbox")

	c.component.El.AddAttribute("data-gomix-component", "SwitchCmp")
	c.component.El.AddAttribute("class", "chakra-switch-cmp")

	c.component.Children(children...)
	return c
}
