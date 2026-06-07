package media_and_icons

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type image struct {
	component *components.Comp
}

func (c *image) Component() *components.Comp {
	return c.component
}

func (c *image) Element() *element.HtmlElement {
	return c.component.El
}

func (c *image) Children(children ...components.IsComponent) *image {
	c.component.Children(children...)
	return c
}

func Image(children ...components.IsComponent) *image {
	c := &image{
		component: &components.Comp{
			El: element.Img(element.Text("")).Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Image")
	c.component.El.AddAttribute("class", "chakra-image")

	c.component.Children(children...)
	return c
}
