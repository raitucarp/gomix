package buttons

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type iconButton struct {
	component *components.Comp
}

func (c *iconButton) Component() *components.Comp {
	return c.component
}

func (c *iconButton) Element() *element.HtmlElement {
	return c.component.El
}

func (c *iconButton) Children(children ...components.IsComponent) *iconButton {
	c.component.Children(children...)
	return c
}

func IconButton(children ...components.IsComponent) *iconButton {
	c := &iconButton{
		component: &components.Comp{
			El: element.Button().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "IconButton")
	c.component.El.AddAttribute("class", "chakra-icon-button")

	c.component.Children(children...)
	return c
}
