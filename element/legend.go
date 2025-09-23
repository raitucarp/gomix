package element

import (
	"golang.org/x/net/html"
)

type legend struct {
	el *HtmlElement
}

func (l *legend) Element() *HtmlElement { return l.el }
func (l *legend) Render() string        { return l.el.Render() }

func Legend(children IsElement) *legend {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "legend",
		},
	}

	el.Children(children)

	l := &legend{el: el}

	return l
}
