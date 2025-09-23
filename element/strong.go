package element

import (
	"golang.org/x/net/html"
)

type strong struct {
	el *HtmlElement
}

func (s *strong) Element() *HtmlElement { return s.el }
func (s *strong) Render() string        { return s.el.Render() }

func Strong(children ...IsElement) *strong {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "strong",
		},
	}

	el.Children(children...)
	s := &strong{el: el}

	return s
}
