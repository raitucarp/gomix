package element

import (
	"golang.org/x/net/html"
)

type s struct {
	el *HtmlElement
}

func (t *s) Element() *HtmlElement { return t.el }
func (t *s) Render() string        { return t.el.Render() }

func S(children IsElement) *s {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "q",
		},
	}

	el.Children(children)
	t := &s{el: el}

	return t
}
