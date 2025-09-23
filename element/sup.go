package element

import (
	"golang.org/x/net/html"
)

type sup struct {
	el *HtmlElement
}

func (s *sup) Element() *HtmlElement { return s.el }
func (s *sup) Render() string        { return s.el.Render() }

func Sup(children ...IsElement) *sup {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "sup",
		},
	}

	el.Children(children...)
	s := &sup{el: el}

	return s
}
