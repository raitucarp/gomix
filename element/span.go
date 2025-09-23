package element

import (
	"golang.org/x/net/html"
)

type span struct {
	el *HtmlElement
}

func (s *span) Element() *HtmlElement { return s.el }
func (s *span) Render() string        { return s.el.Render() }

func Span(children ...IsElement) *span {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "span",
		},
	}

	el.Children(children...)
	s := &span{el: el}

	return s
}
