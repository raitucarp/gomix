package element

import (
	"golang.org/x/net/html"
)

type td struct {
	el *HtmlElement
}

func (t *td) Element() *HtmlElement { return t.el }
func (t *td) Render() string        { return t.el.Render() }

func Td(children ...IsElement) *td {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "td",
		},
	}

	el.Children(children...)
	t := &td{el: el}

	return t
}
