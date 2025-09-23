package element

import (
	"golang.org/x/net/html"
)

type tbody struct {
	el *HtmlElement
}

func (t *tbody) Element() *HtmlElement { return t.el }
func (t *tbody) Render() string        { return t.el.Render() }

func TBody(children ...IsElement) *tbody {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "tbody",
		},
	}

	el.Children(children...)
	t := &tbody{el: el}

	return t
}
