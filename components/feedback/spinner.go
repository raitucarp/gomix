package feedback

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type spinner struct {
	component *components.Comp
}

func (c *spinner) Component() *components.Comp {
	return c.component
}

func (c *spinner) Element() *element.HtmlElement {
	return c.component.El
}

func (c *spinner) Children(children ...components.IsComponent) *spinner {
	c.component.Children(children...)
	return c
}

func Spinner(children ...components.IsComponent) *spinner {
	c := &spinner{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Spinner")
	c.component.El.AddAttribute("class", "chakra-spinner")

	c.component.Children(children...)
	return c
}
