package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type input struct {
	component *components.Comp
}

func (c *input) Component() *components.Comp {
	return c.component
}

func (c *input) Element() *element.HtmlElement {
	return c.component.El
}

func (c *input) Children(children ...components.IsComponent) *input {
	c.component.Children(children...)
	return c
}

func Input(children ...components.IsComponent) *input {
	c := &input{
		component: &components.Comp{
			El: element.Input().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Input")
	c.component.El.AddAttribute("class", "chakra-input")

	c.component.Children(children...)
	return c
}

func (c *input) Variant(v string) *input {
	c.component.El.AddAttribute("data-variant", v)
	return c
}

func (c *input) Size(s string) *input {
	c.component.El.AddAttribute("data-size", s)
	return c
}

func (c *input) IsInvalid(invalid bool) *input {
	if invalid {
		c.component.El.AddAttribute("data-invalid", "")
		c.component.El.AddAttribute("aria-invalid", "true")
	}
	return c
}

func (c *input) IsRequired(req bool) *input {
	if req {
		c.component.El.AddAttributeBool("required")
		c.component.El.AddAttribute("aria-required", "true")
	}
	return c
}
