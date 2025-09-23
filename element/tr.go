package element

import (
	"golang.org/x/net/html"
)

type tr struct {
	el *HtmlElement
}

func (t *tr) Element() *HtmlElement { return t.el }
func (t *tr) Render() string        { return t.el.Render() }

func Tr(children ...IsElement) *tr {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "tr",
		},
	}

	el.Children(children...)
	t := &tr{el: el}

	return t
}
