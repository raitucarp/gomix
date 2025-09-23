package element

import (
	"golang.org/x/net/html"
)

type p struct {
	el *HtmlElement
}

func (t *p) Element() *HtmlElement { return t.el }
func (t *p) Render() string        { return t.el.Render() }

func P(children ...IsElement) *p {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "p",
		},
	}

	el.Children(children...)

	t := &p{el: el}

	return t
}
