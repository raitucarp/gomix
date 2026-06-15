package forms

import (
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type checkbox struct {
	component *components.Comp
}

func (c *checkbox) Component() *components.Comp {
	return c.component
}

func (c *checkbox) Element() *element.HtmlElement {
	return c.component.El
}

func (c *checkbox) Children(children ...components.IsComponent) *checkbox {
	c.component.Children(children...)
	return c
}

func Checkbox(children ...components.IsComponent) *checkbox {
	c := &checkbox{
		component: &components.Comp{
			El: element.Input().Element(),
		},
	}

	c.component.El.AddAttribute("data-gomix-component", "Checkbox")
	c.component.El.AddAttribute("class", "chakra-checkbox")
	c.component.El.AddAttribute("type", "checkbox")

	c.component.Children(children...)
	return c
}

func (c *checkbox) ColorScheme(color string) *checkbox {
	c.component.El.AddAttribute("data-color-scheme", color)
	return c
}
func (c *checkbox) Size(s string) *checkbox {
	c.component.El.AddAttribute("data-size", s)
	return c
}
func (c *checkbox) IsInvalid(invalid bool) *checkbox {
	if invalid {
		c.component.El.AddAttribute("data-invalid", "")
	}
	return c
}
