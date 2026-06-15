package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type textarea struct {
	component *components.Comp
}

func (c *textarea) Component() *components.Comp {
	return c.component
}

func (c *textarea) Element() *element.HtmlElement {
	return c.component.El
}

func (c *textarea) Children(children ...components.IsComponent) *textarea {
	c.component.Children(children...)
	return c
}

func Textarea(children ...components.IsComponent) *textarea {
	c := &textarea{
		component: &components.Comp{
			El: element.TextArea().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Textarea")
	c.component.El.AddAttribute("class", "chakra-textarea")

	c.component.Children(children...)
	return c
}
