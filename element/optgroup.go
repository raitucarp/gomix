package element

import (
	"golang.org/x/net/html"
)

type optgroup struct {
	el *HtmlElement
}

func (o *optgroup) Disabled() *optgroup {
	o.el.AddAttributeBool("disabled")
	return o
}

func (o *optgroup) Label(text string) *optgroup {
	o.el.AddAttribute("label", text)
	return o
}

func (o *optgroup) Element() *HtmlElement { return o.el }
func (o *optgroup) Render() string        { return o.el.Render() }

func OptGroup(children ...IsElement) *optgroup {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "optgroup",
		},
	}

	el.Children(children...)

	o := &optgroup{el: el}

	return o
}
