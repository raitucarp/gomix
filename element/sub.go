package element

import (
	"golang.org/x/net/html"
)

type sub struct {
	el *HtmlElement
}

func (s *sub) Element() *HtmlElement { return s.el }
func (s *sub) Render() string        { return s.el.Render() }

func Sub(children ...IsElement) *sub {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "sub",
		},
	}

	el.Children(children...)
	s := &sub{el: el}

	return s
}
