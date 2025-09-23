package element

import (
	"golang.org/x/net/html"
)

type summary struct {
	el *HtmlElement
}

func (s *summary) Element() *HtmlElement { return s.el }
func (s *summary) Render() string        { return s.el.Render() }

func Summary(children ...IsElement) *summary {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "summary",
		},
	}

	el.Children(children...)
	s := &summary{el: el}

	return s
}
