package typography

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type kbd struct {
	component *components.Comp
}

func (c *kbd) Component() *components.Comp {
	return c.component
}

func (c *kbd) Element() *element.HtmlElement {
	return c.component.El
}

func (c *kbd) Children(children ...components.IsComponent) *kbd {
	c.component.Children(children...)
	return c
}

func Kbd(children ...components.IsComponent) *kbd {
	c := &kbd{
		component: &components.Comp{
			El: element.Kbd(element.Text("")).Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Kbd")
	c.component.El.AddAttribute("class", "chakra-kbd")

	c.component.Children(children...)
	return c
}
