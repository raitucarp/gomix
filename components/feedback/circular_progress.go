package feedback

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type circularProgress struct {
	component *components.Comp
}

func (c *circularProgress) Component() *components.Comp {
	return c.component
}

func (c *circularProgress) Element() *element.HtmlElement {
	return c.component.El
}

func (c *circularProgress) Children(children ...components.IsComponent) *circularProgress {
	c.component.Children(children...)
	return c
}

func CircularProgress(children ...components.IsComponent) *circularProgress {
	c := &circularProgress{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "CircularProgress")
	c.component.El.AddAttribute("class", "chakra-circular-progress")

	c.component.Children(children...)
	return c
}
