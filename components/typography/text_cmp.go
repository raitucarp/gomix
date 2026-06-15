package typography

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type textCmp struct {
	component *components.Comp
}

func (c *textCmp) Component() *components.Comp {
	return c.component
}

func (c *textCmp) Element() *element.HtmlElement {
	return c.component.El
}

func (c *textCmp) Children(children ...components.IsComponent) *textCmp {
	c.component.Children(children...)
	return c
}

func TextCmp(children ...components.IsComponent) *textCmp {
	c := &textCmp{
		component: &components.Comp{
			El: element.P().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "TextCmp")
	c.component.El.AddAttribute("class", "chakra-text-cmp")

	c.component.Children(children...)
	return c
}
