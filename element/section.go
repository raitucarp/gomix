package element

import (
	"golang.org/x/net/html"
)

type section struct {
	el *HtmlElement
}

func (s *section) Element() *HtmlElement { return s.el }
func (s *section) Render() string        { return s.el.Render() }

func Section(children ...IsElement) *section {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "section",
		},
	}

	el.Children(children...)
	s := &section{el: el}

	return s
}
