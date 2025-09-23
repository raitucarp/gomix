package element

import (
	"golang.org/x/net/html"
)

type menu struct {
	el *HtmlElement
}

func (m *menu) Element() *HtmlElement { return m.el }
func (m *menu) Render() string        { return m.el.Render() }

func Menu(children IsElement) *menu {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "menu",
		},
	}

	el.Children(children)

	m := &menu{el: el}

	return m
}
