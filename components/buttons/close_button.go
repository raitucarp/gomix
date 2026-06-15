package buttons

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type closeButton struct {
	component *components.Comp
}

func (c *closeButton) Component() *components.Comp {
	return c.component
}

func (c *closeButton) Element() *element.HtmlElement {
	return c.component.El
}

func (c *closeButton) Children(children ...components.IsComponent) *closeButton {
	c.component.Children(children...)
	return c
}

func CloseButton(children ...components.IsComponent) *closeButton {
	c := &closeButton{
		component: &components.Comp{
			El: element.Button().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "CloseButton")
	c.component.El.AddAttribute("class", "chakra-close-button")

	c.component.Children(children...)
	return c
}
