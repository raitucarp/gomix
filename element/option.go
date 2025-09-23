package element

import (
	"golang.org/x/net/html"
)

type option struct {
	el *HtmlElement
}

func (o *option) Disabled() *option {
	o.el.AddAttributeBool("disabled")
	return o
}

func (o *option) Label(text string) *option {
	o.el.AddAttribute("label", text)
	return o
}

func (o *option) Selected() *option {
	o.el.AddAttributeBool("selected")
	return o
}

func (o *option) Value(value string) *option {
	o.el.AddAttribute("value", value)
	return o
}

func (o *option) Element() *HtmlElement { return o.el }
func (o *option) Render() string        { return o.el.Render() }

func Option(children ...IsElement) *option {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "option",
		},
	}

	el.Children(children...)

	o := &option{el: el}

	return o
}
