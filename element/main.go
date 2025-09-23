package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type main struct {
	el *HtmlElement
}

func (m *main) Value(value int) *main {
	m.el.AddAttribute("value", strconv.Itoa(value))
	return m
}
func (m *main) Element() *HtmlElement { return m.el }
func (m *main) Render() string        { return m.el.Render() }

func Main(children ...IsElement) *main {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "main",
		},
	}

	el.Children(children...)

	m := &main{el: el}

	return m
}
