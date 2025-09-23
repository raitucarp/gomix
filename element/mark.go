package element

import (
	"golang.org/x/net/html"
)

type mark struct {
	el *HtmlElement
}

func (m *mark) Element() *HtmlElement { return m.el }
func (m *mark) Render() string        { return m.el.Render() }

func Mark(children IsElement) *mark {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "mark",
		},
	}

	el.Children(children)

	m := &mark{el: el}

	return m
}
