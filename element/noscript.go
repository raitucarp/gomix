package element

import (
	"golang.org/x/net/html"
)

type noscript struct {
	el *HtmlElement
}

func (n *noscript) Element() *HtmlElement { return n.el }
func (n *noscript) Render() string        { return n.el.Render() }

func NoScript(children IsElement) *noscript {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "noscript",
		},
	}

	el.Children(children)

	n := &noscript{el: el}

	return n
}
