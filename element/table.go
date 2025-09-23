package element

import (
	"golang.org/x/net/html"
)

type table struct {
	el *HtmlElement
}

func (t *table) Element() *HtmlElement { return t.el }
func (t *table) Render() string        { return t.el.Render() }

func Table(children ...IsElement) *table {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "table",
		},
	}

	el.Children(children...)
	t := &table{el: el}

	return t
}
