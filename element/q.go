package element

import (
	"golang.org/x/net/html"
)

type q struct {
	el *HtmlElement
}

func (t *q) Element() *HtmlElement { return t.el }
func (t *q) Render() string        { return t.el.Render() }

func Q(children IsElement) *q {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "q",
		},
	}

	el.Children(children)
	t := &q{el: el}

	return t
}
