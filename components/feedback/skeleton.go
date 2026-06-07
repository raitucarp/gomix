package feedback

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type skeleton struct {
	component *components.Comp
}

func (c *skeleton) Component() *components.Comp {
	return c.component
}

func (c *skeleton) Element() *element.HtmlElement {
	return c.component.El
}

func (c *skeleton) Children(children ...components.IsComponent) *skeleton {
	c.component.Children(children...)
	return c
}

func Skeleton(children ...components.IsComponent) *skeleton {
	c := &skeleton{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Skeleton")
	c.component.El.AddAttribute("class", "chakra-skeleton")

	c.component.Children(children...)
	return c
}
