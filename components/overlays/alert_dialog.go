package overlays

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type alertDialog struct {
	component *components.Comp
}

func (c *alertDialog) Component() *components.Comp {
	return c.component
}

func (c *alertDialog) Element() *element.HtmlElement {
	return c.component.El
}

func (c *alertDialog) Children(children ...components.IsComponent) *alertDialog {
	c.component.Children(children...)
	return c
}

func AlertDialog(children ...components.IsComponent) *alertDialog {
	c := &alertDialog{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "AlertDialog")
	c.component.El.AddAttribute("class", "chakra-alert-dialog")

	c.component.Children(children...)
	return c
}
