package element

import (
	"golang.org/x/net/html"
)

type label struct {
	el *HtmlElement
}

func (l *label) For(elementId string) *label {
	l.el.AddAttribute("for", elementId)
	return l
}

func (l *label) Form(formId string) *label {
	l.el.AddAttribute("form", formId)
	return l
}

func (l *label) Element() *HtmlElement { return l.el }
func (l *label) Render() string        { return l.el.Render() }

func Label(children IsElement) *label {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "label",
		},
	}

	el.Children(children)

	l := &label{el: el}

	return l
}
