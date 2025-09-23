package element

import (
	"golang.org/x/net/html"
)

type mapElement struct {
	el *HtmlElement
}

func (m *mapElement) Name(name string) *mapElement {
	m.el.AddAttribute("name", name)
	return m
}
func (m *mapElement) Element() *HtmlElement { return m.el }
func (m *mapElement) Render() string        { return m.el.Render() }

func Map(children ...IsElement) *mapElement {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "map",
		},
	}

	el.Children(children...)

	m := &mapElement{el: el}

	return m
}
