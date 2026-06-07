package media_and_icons

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type avatar struct {
	component *components.Comp
}

func (c *avatar) Component() *components.Comp {
	return c.component
}

func (c *avatar) Element() *element.HtmlElement {
	return c.component.El
}

func (c *avatar) Children(children ...components.IsComponent) *avatar {
	c.component.Children(children...)
	return c
}

func Avatar(children ...components.IsComponent) *avatar {
	c := &avatar{
		component: &components.Comp{
			El: element.Img(element.Text("")).Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Avatar")
	c.component.El.AddAttribute("class", "chakra-avatar")

	c.component.Children(children...)
	return c
}
