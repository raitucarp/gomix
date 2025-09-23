package element

import (
	"golang.org/x/net/html"
)

type small struct {
	el *HtmlElement
}

func (s *small) Element() *HtmlElement { return s.el }
func (s *small) Render() string        { return s.el.Render() }

func Small(children ...IsElement) *small {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "small",
		},
	}

	el.Children(children...)
	s := &small{el: el}

	return s
}
