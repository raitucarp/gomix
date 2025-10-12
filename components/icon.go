package components

import (
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
	"github.com/raitucarp/gomix/value"
)

type icon struct {
	component *component
}

func (i *icon) Component() *component {
	return i.component
}

func (i *icon) Small() *icon {
	i.component.SizeBy(value.Property("text-sm"))
	return i
}

func (i *icon) Medium() *icon {
	i.component.SizeBy(value.Property("text-md"))
	return i
}

func (i *icon) Large() *icon {
	i.component.SizeBy(value.Property("text-lg"))
	return i
}

func (i *icon) Xl() *icon {
	i.component.SizeBy(value.Property("text-xl"))
	return i
}

func (i *icon) Size2Xl() *icon {
	i.component.SizeBy(value.Property("text-2xl"))
	return i
}

func (i *icon) Size3Xl() *icon {
	i.component.SizeBy(value.Property("text-3xl"))
	return i
}

func (i *icon) Element() *element.HtmlElement {
	return i.component.el
}

func Icon(ic icons.IsIcon) *icon {
	i := &icon{
		component: &component{
			el: ic.Element(),
		},
	}
	i.Large()
	return i
}
