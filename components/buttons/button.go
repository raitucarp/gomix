package buttons

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type button struct {
	component *components.Comp
}

func (c *button) Component() *components.Comp {
	return c.component
}

func (c *button) Element() *element.HtmlElement {
	return c.component.El
}

func (c *button) Children(children ...components.IsComponent) *button {
	c.component.Children(children...)
	return c
}

func Button(children ...components.IsComponent) *button {
	c := &button{
		component: &components.Comp{
			El: element.Button().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Button")
	c.component.El.AddAttribute("class", "chakra-button")

	c.component.Children(children...)
	return c
}

func (c *button) Variant(v string) *button {
	c.component.El.AddAttribute("data-variant", v)
	return c
}

func (c *button) Size(s string) *button {
	c.component.El.AddAttribute("data-size", s)
	return c
}

func (c *button) ColorScheme(color string) *button {
	c.component.El.AddAttribute("data-color-scheme", color)
	return c
}

func (c *button) IsLoading(loading bool) *button {
	if loading {
		c.component.El.AddAttribute("data-loading", "")
		c.component.El.AddAttributeBool("disabled")
	}
	return c
}
