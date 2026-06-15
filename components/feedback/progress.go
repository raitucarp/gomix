package feedback

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type progress struct {
	component *components.Comp
}

func (c *progress) Component() *components.Comp {
	return c.component
}

func (c *progress) Element() *element.HtmlElement {
	return c.component.El
}

func (c *progress) Children(children ...components.IsComponent) *progress {
	c.component.Children(children...)
	return c
}

func Progress(children ...components.IsComponent) *progress {
	c := &progress{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Progress")
	c.component.El.AddAttribute("class", "chakra-progress")

	c.component.Children(children...)
	return c
}
