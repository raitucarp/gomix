package element

import (
	"golang.org/x/net/html"
)

type search struct {
	el *HtmlElement
}

func (s *search) Element() *HtmlElement { return s.el }
func (s *search) Render() string        { return s.el.Render() }

func Search(children ...IsElement) *search {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "search",
		},
	}

	el.Children(children...)
	s := &search{el: el}

	return s
}
