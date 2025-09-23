package element

import (
	"golang.org/x/net/html"
)

type samp struct {
	el *HtmlElement
}

func (s *samp) Element() *HtmlElement { return s.el }
func (s *samp) Render() string        { return s.el.Render() }

func Samp(children ...IsElement) *samp {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "samp",
		},
	}

	el.Children(children...)
	s := &samp{el: el}

	return s
}
