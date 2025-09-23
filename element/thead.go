package element

import (
	"golang.org/x/net/html"
)

type thead struct {
	el *HtmlElement
}

func (t *thead) Element() *HtmlElement { return t.el }
func (t *thead) Render() string        { return t.el.Render() }

func THead(children ...IsElement) *thead {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "thead",
		},
	}

	el.Children(children...)
	t := &thead{el: el}

	return t
}
