package element

import (
	"golang.org/x/net/html"
)

type output struct {
	el *HtmlElement
}

func (o *output) For(elementId string) *output {
	o.el.AddAttribute("for", elementId)
	return o
}

func (o *output) Form(formId string) *output {
	o.el.AddAttribute("form", formId)
	return o
}

func (o *output) Name(name string) *output {
	o.el.AddAttribute("name", name)
	return o
}

func (o *output) Element() *HtmlElement { return o.el }
func (o *output) Render() string        { return o.el.Render() }

func Output(children ...IsElement) *output {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "output",
		},
	}

	el.Children(children...)

	o := &output{el: el}

	return o
}
